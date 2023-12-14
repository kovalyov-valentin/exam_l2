package pkg

// Сборщик для компьютеров lenovo 
type LenovoCollector struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	GPU     int
}

func (lc *LenovoCollector) SetCore() {
	lc.Core = 4
}

func (lc *LenovoCollector) SetBrand() {
	lc.Brand = "Lenovo"
}

func (lc *LenovoCollector) SetMemory() {
	lc.Memory = 32
}

func (lc *LenovoCollector) SetMonitor() {
	lc.Monitor = 2
}

func (lc *LenovoCollector) SetGPU() {
	lc.GPU = 2
}

func (lc *LenovoCollector) GetComputer() Computer {
	return Computer{
		Core: lc.Core,
		Brand: lc.Brand,
		Memory: lc.Memory,
		Monitor: lc.Monitor,
		GPU: lc.GPU,
	}
}
