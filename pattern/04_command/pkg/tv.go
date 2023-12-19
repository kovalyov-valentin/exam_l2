package pkg

import "fmt"

type tv struct {
	isRunning bool
}

func NewTv() *tv {
	return &tv{}
}

func (t *tv) On() {
	t.isRunning = true
	fmt.Println("Телевизор включен")
}

func (t *tv) Off() {
	t.isRunning = false
	fmt.Println("Телевизор выключен")
}
