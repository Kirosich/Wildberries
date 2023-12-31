package Finder

import (
	"GrepTask/Flags"
	"GrepTask/helpers"
	"fmt"
	"regexp"
	"strings"
)

func Find(flags *Flags.Flags) {
	lines := Flags.OpenFile(flags)
	//Проверка на присутствие флагов
	IsFlags := false

	// Массив с результатом
	var res [][]string

	if flags.A != "" {
		IsFlags = true
		res = Flags.FlagA(lines, flags)
	}

	if flags.B != "" {
		IsFlags = true
		res = Flags.FlagB(lines, flags)
	}

	if flags.C != "" {
		IsFlags = true
		res = Flags.FlagC(lines, flags)
	}

	if *flags.Count {
		IsFlags = true
		res = Flags.FlagCount(lines, flags)
	}

	if *flags.N {
		IsFlags = true
		res = Flags.FlagN(lines, flags)
	}

	if *flags.F {
		IsFlags = true
		res = Flags.FlagF(lines, flags)
	}

	if *flags.V {
		IsFlags = true
		res = Flags.FlagV(lines, flags)
	}

	if !IsFlags {
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
				matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
				helpers.HandleError(err)
				if matched {
					res = append(res, lines[i])
					break
				}
			}
		}
	}

	PrintLines(res)

}

func PrintLines(res [][]string) {
	for i := 0; i < len(res); i++ {
		if i == len(res)-1 && strings.Join(res[i], "") == "--" {
			break
		}
		fmt.Println(strings.Join(res[i], " "))
	}
}
