package pkg

import "fmt"

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

// Вот реализация фабричного метода. Будет инициализировать структуры и возвращать интерфейс Computer
// Централизованный конструктор создания объектов
func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s Несуществующий тип объекта\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}
}
