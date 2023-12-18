package main

import "github.com/kovalyov-valentin/exam_l2/pattern/05_chain_of_resp/pkg"

func main() {
	device := &pkg.Device{Name: "Device-1"}                   // Создание устройства
	updateService := &pkg.UpdateDataService{Name: "Update-1"} // Создание сервиса обновления данных
	dataService := &pkg.DataService{}                         // Создание сервиса сохранения данных

	// Теперь выстраиваем цепочку, что за чем вызывать

	device.SetNext(updateService)      // Устройство передает данные в сервис обновления данных
	updateService.SetNext(dataService) // Сервис обновления передает данные сервису сохранения

	data := &pkg.Data{}  // Создаем данные
	device.Execute(data) // У самого первого звена вызывает обработку данных
}
