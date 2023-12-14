package pkg

// Наши типы
const (
	AcerCollectorType = "acer"
	LenovoCollectorType = "lenovo"
)


// Интерфейс нашего сборщика
type Collector interface {
	SetCore() // Ядра
	SetBrand() // Бренд
	SetMemory() // ОЗУ
	SetMonitor() // Монитор
	SetGPU() // Видеокарта
	GetComputer() Computer // Метод для возвращения укомплектованного компьютера
}

// Функция возвращающая нашего сборщика, в зависимости от того аргумента, который мы передадим
func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil 
	case AcerCollectorType:
		return &AcerCollector{}
	case LenovoCollectorType:
		return &LenovoCollector{ }
	}
}
