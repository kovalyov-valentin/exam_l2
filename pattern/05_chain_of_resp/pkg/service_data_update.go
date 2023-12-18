package pkg

import "fmt"

// 2.  Обновление данных которые нам передало устройство. То есть обработка данных
type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource { // Проверка состояния что входящие данные уже были обработаны
		fmt.Printf("Data in service [%s] is already update.\n", upd.Name)
		upd.Next.Execute(data)
		return
	}

	fmt.Printf("Update data from service [%s]\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}
