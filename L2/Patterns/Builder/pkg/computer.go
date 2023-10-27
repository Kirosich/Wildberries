package pkg

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     int
}

func (pc *Computer) Print() {
	fmt.Printf("%s Core [%d] Mem:[%d] GraphicCard:[%d] Monitor: [%d] \n", pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}
