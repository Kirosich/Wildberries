package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	// Преобразуем строку в массив символов
	slice := strings.Split(s, "")
	// Сортируем массив символов
	sort.Strings(slice)
	// Собираем отсортированный массив обратно в строку
	return strings.Join(slice, "")
}

func isInSlice(s string, slice []string) bool {
	for _, val := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func Anagrams(mas *[]string) *map[string]*[]string {
	// один элемент скип
	anagramMap := make(map[string]*[]string) // Ключ - первое попавшеесе слово, Значение ссылка на массив множества

	var wordsSlice []string // Массив слов анаграмм

	length := len(*mas)
	for i := 0; i < length; i++ {

		wordsSlice = []string{}
		wordsSlice = append(wordsSlice, (*mas)[i]) // Добавляем первое слово

		for j := 0; j < length; j++ {
			_, ok := anagramMap[(*mas)[j]] // Проверяем есть ли это слово уже как ключ
			if (*mas)[i] == (*mas)[j] {
				continue
			}
			if sortString((*mas)[i]) == sortString((*mas)[j]) {
				if ok || isInSlice((*mas)[j], wordsSlice) {
					break
				}
				wordsSlice = append(wordsSlice, (*mas)[j])
			}

		}
		if len(wordsSlice) != 1 {
			copiedSlice := make([]string, len(wordsSlice))
			copy(copiedSlice, wordsSlice)
			sort.Strings(copiedSlice)
			anagramMap[wordsSlice[0]] = &copiedSlice
		}
	}
	return &anagramMap
}

func main() {
	mas := []string{"пятак", "Тяпка", "Пятка", "Листок", "слиток", "арап", "столик", "п", "да", "ад", "огурец", "неземной", "гной", "ключ", "пара", "да", "ад", "карта", "карат", "катар", "кулон", "уклон", "колун", "клоун"}

	for i, val := range mas {
		mas[i] = strings.ToLower(val)
	}

	anagramMap := Anagrams(&mas)
	fmt.Println(anagramMap)
}
