package gows

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

/**
  @Author   : bob
  @Datetime : 2023-05-21 下午 5:55
  @File     : client_manager.go
  @Desc     :
*/

type ClientManager struct {
	Clients     map[*Client]bool
	ClientsLock sync.RWMutex
	Register    chan *Client
	Unregister  chan *Client
	Broadcast   chan []byte
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

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.EventRegister(conn)

		}

	}

}
