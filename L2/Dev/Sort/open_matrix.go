package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func OpenMatrix() [][]string {
	// Открываем файл
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return nil
	}
	defer file.Close()

	var matrix [][]string

	// Создаем сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Разбиваем строку на слова
		words := strings.Fields(scanner.Text())
		// Добавляем слова в двумерный массив
		matrix = append(matrix, words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
		return nil
	}

	// Выводим содержимое двумерного массива
	return matrix
}
