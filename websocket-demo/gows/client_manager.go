package gows

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

/**
  @Author   : bob
  @Datetime : 2023-05-21 下午 5:55
  @File     : client_manager.go
  @Desc     :
*/

// 多协程读写map，需要加锁。
// Clients
type ClientManager struct {
	Clients     map[*Client]bool
	ClientsLock sync.RWMutex
	Register    chan *Client
	Unregister  chan *Client
	Broadcast   chan []byte
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
	return clientManager
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			log.Info("get client register event from channel")
			manager.EventRegister(conn)
		case conn := <-manager.Unregister:
			log.Info("get client unregister event from channel")
			manager.EventUnregister(conn)
		case msg := <-manager.Broadcast:
			log.Info("get broadcast msg from channel")
			clients := manager.GetClients()
			for conn := range clients {
				select {
				case conn.Send <- msg:
				default:
					close(conn.Send)
				}
			}
		}
	}
}

func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client] = true
}

func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	if _, ok := manager.Clients[client]; ok {
		delete(manager.Clients, client)
		log.Info("delete client success")
	}
}

func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
	log.Info("register client success")
}

func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)
	close(client.Send)
	log.Info("unregister client success")
}

func (manager *ClientManager) GetClients() (clients map[*Client]bool) {
	clients = make(map[*Client]bool)
	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value
		return true
	})
	return
}

func (manager *ClientManager) GetClientsLen() int {
	len := len(manager.Clients)
	return len
}

func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	for k, v := range manager.Clients {
		result := f(k, v)
		if result == false {
			return
		}
	}
	return
}

func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	timeout = false
	return timeout
}

func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())
	clients := clientManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			log.Warn("Connection %s is timeout, close it", client)
			client.Socket.Close()
		}
	}

}

func GetManagerInfo(isDebug string) (managerInfo map[string]interface{}) {
	namagerInfo = make(map[string]interface{})
	managerInfo["clientsLen"] = clientManager.GetClientsLen()

	return managerInfo

}
