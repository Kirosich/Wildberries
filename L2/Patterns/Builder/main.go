package main

import "Builder/pkg"

func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")

	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
