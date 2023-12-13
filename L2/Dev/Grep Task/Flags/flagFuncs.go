package Flags

import (
	"GrepTask/helpers"
	"regexp"
	"strconv"
)

func FlagA(lines [][]string, flags *Flags) [][]string {

	var res [][]string
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
			helpers.HandleError(err)
			if matched {

				res = append(res, lines[i])

				AfterLines, _ := strconv.Atoi(flags.A)
				AfterLinesI := i + AfterLines

				for ln := i; ln < AfterLinesI; ln++ {
					if i+AfterLines-1 > len(lines) {
						break
					}
					i++
					res = append(res, lines[i])
				}

				res = append(res, []string{"--"})
				break
			}
		}
	}
	return res
}

func FlagB(lines [][]string, flags *Flags) [][]string {

	var res [][]string
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
			helpers.HandleError(err)
			if matched {
				BeforeLines, _ := strconv.Atoi(flags.B)
				BeforeLinesI := i - BeforeLines

				for ln := BeforeLinesI; ln < i; ln++ {
					if BeforeLinesI < 0 {
						BeforeLinesI++
						continue
					}
					res = append(res, lines[ln])
				}

				res = append(res, lines[i])

				res = append(res, []string{"--"})
				break
			}
		}
	}
	return res
}

func FlagC(lines [][]string, flags *Flags) [][]string {

	var res [][]string
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
			helpers.HandleError(err)
			if matched {
				BeforeLines, _ := strconv.Atoi(flags.C)
				BeforeLinesI := i - BeforeLines

				for ln := BeforeLinesI; ln < i; ln++ {
					if BeforeLinesI < 0 {
						BeforeLinesI++
						continue
					}

					res = append(res, lines[ln])
				}

				res = append(res, lines[i])

				AfterLines, _ := strconv.Atoi(flags.C)
				AfterLinesI := i + AfterLines

				for ln := i; ln < AfterLinesI; ln++ {
					if i+AfterLines-1 > len(lines) {
						break
					}
					i++
					res = append(res, lines[i])
				}

				res = append(res, []string{"--"})
				break
			}
		}
	}
	return res
}
