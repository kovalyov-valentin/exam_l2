package pkg

import "fmt"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Computer {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  16,
		Monitor: true,
	}
}

func (n Notebook) GetType() string {
	return n.Type
}

func (n Notebook) PrintDetails() {
	fmt.Printf("%s Core: %d Memory: %d Monitor: %t\n", n.Type, n.Core, n.Memory, n.Monitor)
}
