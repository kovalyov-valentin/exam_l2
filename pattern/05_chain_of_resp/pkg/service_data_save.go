package pkg

import "fmt"

// 3. Данный сервис отвечает за сохранение обработанных данных
type DataService struct {
	Next Service
}

func (ds *DataService) Execute(data *Data) {
	if !data.UpdateSource { // Если данные не были успешно обработаны, то возвращаем сообщение
		fmt.Println("Data not update")
		return
	}
	fmt.Println("Data save")
}

func (ds *DataService) SetNext(svc Service) {
	ds.Next = svc
}
