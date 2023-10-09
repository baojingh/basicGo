package strategy

import "fmt"

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 10:56
  @File     : strategy.go
  @Desc     :
*/

type StorageStrategy interface {
	Save(name string, data string) error
}

type NormalStorage struct{}
type EncryptStorage struct{}

func (normal NormalStorage) Save(name string, data string) error {
	fmt.Println("normal storage")
	return nil
}

func (normal EncryptStorage) Save(name string, data string) error {
	fmt.Println("encrypt storage")
	return nil
}

var strategy = map[string]StorageStrategy{
	"normal":  &NormalStorage{},
	"encrypt": &EncryptStorage{},
}

func NewStorageStrategy(name string) (StorageStrategy, error) {
	s, ok := strategy[name]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("strategy not found")
}
