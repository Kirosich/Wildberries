package pkg

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     int
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		GraphicCard: collector.GraphicCard,
		Monitor:     collector.Monitor,
	}
}
