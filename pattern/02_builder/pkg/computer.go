package pkg

import "fmt"

// Структура компьютера, который мы будем возвращать
type Computer struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	GPU     int
}

// Функция вывода информации 
func (c *Computer) Print() {
	fmt.Printf("%s Core:[%d] Memory:[%d] Monitor:[%d] GPU:[%d]\n", c.Brand, c.Core, c.Memory, c.Monitor, c.GPU)
}
