package main

import "fmt"

// Интерфейс для клиента
type Target interface {
	Operation()
}

// Клиент
type Client struct {
}

// Клиентская операция
func (c *Client) ClientOperation(adap Target) {
	fmt.Println("Клиент выполняет операцию, во время который необходимо вызвать метод сервиса...")
	adap.Operation()
	fmt.Println("Работа клиента завершена!")
}

// Сервис
type Service struct {
}

// Метод адаптируемого сервиса, который нужно вызвать в клиенте
func (adaptee *Service) AdaptedOperation() {
	fmt.Println("Выполняется сервисный метод...")
}

// Адаптер
type Adapter struct {
	*Service
}

// Реализация метода интерфейса для клиента
func (adapter *Adapter) Operation() {
	adapter.AdaptedOperation()
}

// Конструктор
func NewAdapter(adaptee *Service) Target {
	return &Adapter{adaptee}
}

func main() {
	adapter := NewAdapter(&Service{})
	client := Client{}
	client.ClientOperation(adapter)
}
