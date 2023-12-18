package pkg

import "fmt"

// В этом файле происходит реализация поведения для сервиса получения данных

// 1. Структура устройства, которое передает от источника данных сами данные
type Device struct {
	Name string
	Next Service
}

// Реализация интерфейса нашего сервиса
func (d *Device) Execute(data *Data) {
	if data.GetSource { // Проверка состояния наших данных, т.е. если у нас входные данные уже были обработаны,
		// то мы выводим сообщение о том, что данные уже были ранее обработаны
		fmt.Printf("Data from device [%s] already get.\n", d.Name)
		d.Next.Execute(data) // Если данные уже были обработаны, то мы передаем следующему звену задачу на выполнение
		return
	}
	// Если исходные данные не были приняты, то мы фиксируем это, ставим отметку о том, что данные нужно обработать и передаем следующему звену на обработку
	fmt.Printf("Get data from device [%s].\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

// Задаем следующее звено
func (d *Device) SetNext(svc Service) {
	d.Next = svc

}
