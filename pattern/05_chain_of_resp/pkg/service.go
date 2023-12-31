package pkg

type Service interface {
	Execute(*Data)   // Выполнение операций с полученными данными
	SetNext(Service) // Задача следующего исполнителя. Передаем в параметры следующий сервис, который идет и после этого
}

// Структура наших данных, которые будут переходить от одного сервиса к другому
type Data struct {
	GetSource    bool // Флаг, который является признаком того, приняты данные или нет от источника
	UpdateSource bool // Ставит отметку тот сервис, который обработал данные
}
