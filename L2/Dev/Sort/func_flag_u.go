package main

func FuncFlagU(matrix [][]string, TargetColumn int, all bool) [][]string {
	// Создаём новую матрицу для того, чтобы заппендить столько сколько нужно
	TargetColumn--
	var newMatrix [][]string

	// Создаём срез, в которой будут идти элементы без повтора
	var newSlice []string

	if all {
		for _, val := range matrix {
			newSlice = DeleteRepeating(val)
			newMatrix = append(newMatrix, newSlice)
		}
		return newMatrix

	}

	if !all {
		for index, val := range matrix {
			if index == TargetColumn {
				newSlice = DeleteRepeating(val)
				newMatrix = append(newMatrix, newSlice)
			} else {
				newMatrix = append(newMatrix, val)
			}
		}
		return newMatrix
	}
	return nil
}

func DeleteRepeating(Slice []string) []string { // Функция удаления повторяющихся строк

	// Создаём новый массив, который вернём без повторений
	var newSlice []string

	// Булевая переменная, которая сообщает найдено ли повторение
	var foundRepeat bool

	// Цикл в котором всё и происходит
	for i, val1 := range Slice {
		foundRepeat = false
		for _, val2 := range Slice[i+1:] {
			if val1 == val2 { // Если повторение нашлось
				foundRepeat = true // Просто пропускаем элемент Slice[i] и не добавляем в новый массив
				break
			}
		}
		if !foundRepeat { // Если повторение не нашлось
			newSlice = append(newSlice, Slice[i]) // Добавляем в новый массив
		}
	}

	return newSlice
}
