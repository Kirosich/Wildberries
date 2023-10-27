package pkg

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 8
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "HP"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 2
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		GraphicCard: collector.GraphicCard,
		Monitor:     collector.Monitor,
	}
}
