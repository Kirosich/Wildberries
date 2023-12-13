package helpers

import (
	"fmt"
	"os"
)

func AreSlicesEqual(slice1, slice2 []string) bool {
	// Проверка наличия одинаковой длины
	if len(slice1) != len(slice2) {
		return false
	}

	// Проверка наличия одинаковых элементов
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
func HandleError(err error) {
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		os.Exit(1)
	}
}
