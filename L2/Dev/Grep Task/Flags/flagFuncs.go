package Flags

import (
	"GrepTask/helpers"
	"regexp"
	"strconv"
	"strings"
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
				break
			}
		}
	}
	return res
}

func FlagC(lines [][]string, flags *Flags) [][]string {

	var res [][]string

	// map for [line]index - to not repeat lines

	mapLines := make(map[string]int)

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

					key := strings.Join(lines[ln], " ")

					if _, seen := mapLines[key]; seen {
						if mapLines[key] == ln {
							continue
						} else {
							res = append(res, lines[ln])
							mapLines[strings.Join(lines[ln], " ")] = ln
						}
					} else {
						res = append(res, lines[ln])
						mapLines[strings.Join(lines[ln], " ")] = ln
					}

				}

				key := strings.Join(lines[i], " ")

				if _, seen := mapLines[key]; seen {
					if mapLines[key] == i {
						continue
					} else {
						res = append(res, lines[i])
						mapLines[strings.Join(lines[i], " ")] = i
					}
				} else {
					res = append(res, lines[i])
					mapLines[strings.Join(lines[i], " ")] = i
				}

				AfterLines, _ := strconv.Atoi(flags.C)
				AfterLinesI := i + AfterLines

				for ln := i; ln < AfterLinesI; ln++ {
					if i+AfterLines-1 > len(lines) {
						break
					}
					i++

					key := strings.Join(lines[i], " ")

					if _, seen := mapLines[key]; seen {
						if mapLines[key] == i {
							continue
						} else {
							res = append(res, lines[i])
							mapLines[strings.Join(lines[i], " ")] = i
						}
					} else {
						res = append(res, lines[i])
						mapLines[strings.Join(lines[i], " ")] = i
					}

				}
				break
			}
		}
	}

	return res
}

func FlagCount(lines [][]string, flags *Flags) [][]string {
	var res [][]string
	var counter int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
			helpers.HandleError(err)
			if matched {
				counter++
				break
			}
		}
	}
	res = append(res, []string{strconv.Itoa(counter)})
	return res
}

func FlagN(lines [][]string, flags *Flags) [][]string {
	var res [][]string

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
			helpers.HandleError(err)
			if matched {
				index := strconv.Itoa(i + 1)
				res = append(res, []string{index, strings.Join(lines[i], " ")})
				break
			}
		}
	}

	return res

}

func FlagF(lines [][]string, flags *Flags) [][]string {
	var res [][]string

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			matched, err := regexp.MatchString("^"+flags.TargetWord+"$", lines[i][j])
			helpers.HandleError(err)
			if matched {
				res = append(res, lines[i])
				break
			}
		}
	}
	return res
}

func FlagV(lines [][]string, flags *Flags) [][]string {
	var res [][]string

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			matched, err := regexp.MatchString(flags.TargetWord, lines[i][j])
			helpers.HandleError(err)
			if matched {
				break
			}
			res = append(res, lines[i])
			break
		}
	}

	return res
}
