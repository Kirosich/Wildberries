package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	// Матрица данных
	var dataMatrix [][]string
	dataMatrix = OpenMatrix()
	var Matrix [][]string

	// Переворот матрицы так, чтобы было удобно

	matrixColumns := len(dataMatrix[0]) - 1
	matrixLines := len(dataMatrix) - 1

	for j := 0; j < matrixColumns+1; j++ {
		Slice := []string{} // Каждую итерацию создаём новый слайс для нового столбца
		for i := 0; i < matrixLines+1; i++ {
			Slice = append(Slice, dataMatrix[i][j]) // Добавляем старые значение столбца в отдельный массив
		}
		Matrix = append(Matrix, Slice)
	}

	//Флаги
	var FlagK string
	flag.StringVar(&FlagK, "k", "", "Flag for target column")
	FlagU := flag.Bool("u", false, "Flag for deleting repeating strings")
	FlagN := flag.Bool("n", false, "Flag for sorting numbers in increase")
	FlagR := flag.Bool("r", false, "Flag for sorting number in decrease")

	flag.Parse() // Парсинг флагов

	// Переменная все ли столбцы надо обработать
	var all bool

	//Колонка, которую выбрал пользователь
	var TargetColumn int
	if FlagK != "" {
		//Определенный столбец, поэтому
		all = false

		//Колонка, которую выбрал пользователь
		TargetColumn, _ = strconv.Atoi(FlagK)
		if *FlagU {
			Matrix = FuncFlagU(Matrix, TargetColumn, all)
		}
		if *FlagN || *FlagR {
			Matrix = FuncFlagN(Matrix, TargetColumn, all, *FlagR)
		}
	} else {
		all = true
		if *FlagU {
			Matrix = FuncFlagU(Matrix, TargetColumn, all)
		}

		if *FlagN || *FlagR {
			Matrix = FuncFlagN(Matrix, TargetColumn, all, *FlagR)
		}
	}

	var maxColumn int
	for i := 0; i < len(Matrix); i++ {
		if maxColumn < len(Matrix[i]) {
			maxColumn = len(Matrix[i])
		}
	}

	for columns := 0; columns < maxColumn; columns++ {
		for lines := 0; lines < len(Matrix); lines++ {

			if columns < len(Matrix[lines]) { // columns - какая колонка, если колонка выходит за границы, то ничего не делать
				fmt.Print(Matrix[lines][columns])
				fmt.Print(" ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

}
