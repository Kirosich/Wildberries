package main

import (
	"GrepTask/Flags"
)

// Запуск
func Initialize() *Flags.Flags {
	Flags := Flags.InitFlags()

	return Flags
}
