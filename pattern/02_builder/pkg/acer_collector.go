package pkg

// Сборщик для компьютеров acer 
type AcerCollector struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	GPU     int
}

func (ac *AcerCollector) SetCore() {
	ac.Core = 4
}

func (ac *AcerCollector) SetBrand() {
	ac.Brand = "Acer"
}

func (ac *AcerCollector) SetMemory() {
	ac.Memory = 16
}

func (ac *AcerCollector) SetMonitor() {
	ac.Monitor = 1
}

func (ac *AcerCollector) SetGPU() {
	ac.GPU = 1
}


func (ac *AcerCollector) GetComputer() Computer {
	return Computer{
		Core: ac.Core,
		Brand: ac.Brand,
		Memory: ac.Memory,
		Monitor: ac.Monitor,
		GPU: ac.GPU,
	}
}
