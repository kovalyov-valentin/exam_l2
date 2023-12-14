package pkg

// Структра "завода" по производству компов
type Factory struct {
	Collector Collector
}

// Конструктор для инициализации нового завода
func NewFactory(collector Collector) *Factory {
	return &Factory{
		Collector: collector,
	}
}

// Функция которая позволяет менять комплектацию в любой момент времени.
func (f *Factory) SetCollector(collector Collector) {
	f.Collector = collector
}

// Метод где последовательно выполняются шаги сборки
// Это основная функция которая возвращает укомплектованную сборку на существующем заводе
func (f *Factory) CreateComputer() Computer{
	f.Collector.SetCore()
	f.Collector.SetMemory()
	f.Collector.SetBrand()
	f.Collector.SetGPU()
	f.Collector.SetMonitor()
	return f.Collector.GetComputer()
}

