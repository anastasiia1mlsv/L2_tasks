package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	//список слов для поиска
	lib := []string{"пятак", "пама", "тяпка", "ксам", "пятка", "аапм"}

	//создаем два массива: двумерный массив состоящий из массивов рун
	//и массив контрольных сум для каждого слова
	runes := make([][]rune, 0, len(lib))
	checkSum := make([]int, 0, len(lib))

	//заполняем двумерный массив рун конвертированными значениями из (12)
	//и массив контрольных сумм для каждого слова
	for _, i := range lib {
		runes = append(runes, []rune(i))
		fmt.Println(getWordSum([]rune(i)))
		checkSum = append(checkSum, getWordSum([]rune(i)))
	}

	//переводим результат выполнения в json для удобного чтения
	result, err := json.Marshal(*searchForAnagram(&runes, checkSum))
	if err != nil {
		fmt.Println("Error at marshaling")
		return
	}

	fmt.Println(string(result))

}

// Функция сортирует двумерный массив рун и возвращает мапу строк
func searchForAnagram(words *[][]rune, checkSum []int) *map[string][]string {

	//создаем мапу для сохранения добавленных первично данных в словарь
	//и мапу структур для проверки встречалось ли слово и какой у него глобальный ключ
	result := make(map[string][]string)
	tempSums := make(map[int]struct {
		was     bool
		firstID string
	})

	//итерируемся по длине массива контрольных сумм
	for i := 0; i < len(checkSum); i += 1 {

		//если в структуре для данного слова отрицательный флаг встречи
		if tempSums[checkSum[i]].was {
			//проверяем длину конечного слова чтобы убедится что контрольная сумма не ошибочна
			if len(string((*words)[i])) == len(tempSums[checkSum[i]].firstID) {
				//добавляем в результирующий массив текущее слово по ключу из проверочной мапы
				result[tempSums[checkSum[i]].firstID] = append(result[tempSums[checkSum[i]].firstID], string((*words)[i]))
			}
		} else {
			//иначе инициализируем проверочный массив по данному ключу
			tempSums[checkSum[i]] = struct {
				was     bool
				firstID string
			}{was: true, firstID: string((*words)[i])}
		}
	}

	//итерируясь сортируем анаграммы для каждой ненулевой мапы
	for _, i := range result {
		if i != nil {
			sort.Strings(i)
		}
	}

	return &result
}

// Получаем контрольную сумму для массива рун
func getWordSum(word []rune) int {
	sum := 0
	for _, i := range word {
		sum += int(i)
	}
	return sum
}
