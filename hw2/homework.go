package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	//easy
	fmt.Println("1 easy\n")
	array := []int{1, 2, 3, 4, 5}
	sum := array[0] + array[1] + array[2] + array[3] + array[4]
	fmt.Println(sum)
	//medium
	fmt.Println("2 medium\n")
	names := make([][]string, 5)
	letter := []string{"С"}
	all := []string{"Ваня", "Оля", "Саша", "Маша", "Сережа"}
	result := []string{}
	for i := range names {
		names[i] = make([]string, 5)
		nameid := i % len(all)
		selected := all[nameid]
		names[i][0] = selected
		first := names[i][0][:2]
		if first == letter[0] {
			result = slices.Insert(result, len(result), names[i][0])
		}

	}
	fmt.Println(len(result), result)
	//hard
	fmt.Println("3 hard\n")
	tekst1 := "Входной текст можно представить как строку. Программа должна разбить текст на слова, удалить знаки препинания и привести все слова к нижнему регистру перед подсчетом. Затем программа должна вывести на экран список уникальных слов вместе с количеством их вхождений в текст."
	tekstcls := strings.ReplaceAll(tekst1, ",", "")
	tekst2 := strings.ReplaceAll(tekstcls, ".", "")
	tekst3 := strings.ToLower(tekst2)
	text := strings.Fields(tekst3)
	result1 := CountOnn(text)
	fmt.Println(text, result1)
	// text := make([][]string, len(tekst4))
	// for i := range text {
	// 	selected1 := tekst4[i]
	// 	text[i] = make([]string, 1)
	// 	text[i][0] = selected1
	// }
	// fmt.Println(text)
	// fmt.Println(result)
}
func CountOnn(text []string) map[string]int {
	cnt := make(map[string]int)
	for _, i := range text {
		cnt[i]++
	}
	return cnt
}
