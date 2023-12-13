package Flags

import (
	"flag"
	"os"
)

type Flags struct {
	TargetWord string // TargetWord
	FileName   string // FileName
	A          string // -A | After | Вернуть n строк после targetLine
	B          string // -B | Before | Вернуть n строк перед targetLine
	C          string // -C | Context | Вернуть +-n строк между targetLine
	Count      *bool  // -c | Count | Вернуть кол-во строк с этими словами
	I          *bool  // -i | Ignore Case | Игнорировать регистр при поиске
	V          *bool  // -v | Invert | Вместо совпадения - исключать
	F          *bool  // -F | Fixed | Вместо регулярного выражения, конкретная строка
	N          *bool  // -n | Line num | Вывести номер строки перед самой строкой
}

func InitFlags() *Flags {

	var FlagA, FlagB, FlagC string // Для флагов у которых есть аргумент

	TargetWord := os.Args[len(os.Args)-2]
	FileName := os.Args[len(os.Args)-1]

	flag.StringVar(&FlagA, "A", "", "Вернуть n строк после targetLine")   // FlagA
	flag.StringVar(&FlagB, "B", "", "Вернуть n строк перед targetLine")   // FlagB
	flag.StringVar(&FlagC, "C", "", "Вернуть +-n строк между targetLine") // FlagC
	FlagCount := flag.Bool("c", false, "Вернуть кол-во строк с targetWord")
	FlagI := flag.Bool("i", false, "Игнорировать регистр при поиске")
	FlagV := flag.Bool("v", false, "Вместо совпадения - исключать")
	FlagF := flag.Bool("F", false, "Вместо регулярного выражения, конкретная строка")
	FlagN := flag.Bool("n", false, "Вывести номер строки перед самой строкой")

	flag.Parse() // Считываем флаги

	// Создаём экземпляр флагов
	Flags := &Flags{
		TargetWord: TargetWord,
		FileName:   FileName,
		A:          FlagA,
		B:          FlagB,
		C:          FlagC,
		Count:      FlagCount,
		I:          FlagI,
		V:          FlagV,
		F:          FlagF,
		N:          FlagN,
	}

	return Flags
}
