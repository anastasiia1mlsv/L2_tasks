package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Тестовые строки
	testStrings := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, str := range testStrings {
		result, err := unpack(str)
		if err != nil {
			fmt.Printf("Ошибка при распаковке %q: %v\n", str, err)
		} else {
			fmt.Printf("Распаковано %q: %q\n", str, result)
		}
	}
}

func unpack(input string) (string, error) {
	var result []rune  // Срез для хранения результата
	var prev rune      // Предыдущий символ
	var isEscaped bool // Флаг экранирования

	for _, r := range input {
		if isEscaped {
			// Текущий символ обрабатывается как литерал
			result = append(result, r)
			prev = r
			isEscaped = false
			continue
		}

		if r == '\\' {
			// Начало экранированной последовательности
			isEscaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if prev == 0 {
				// Цифра без предшествующего символа
				return "", fmt.Errorf("некорректная строка: цифра %q без предшествующего символа", r)
			}

			// Повторяем предыдущий символ count - 1 раз
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", fmt.Errorf("некорректная цифра %q: %v", r, err)
			}
			if count == 0 {
				// Удаляем предыдущий символ, если количество равно нулю
				if len(result) > 0 {
					result = result[:len(result)-1]
				}
			} else {
				for i := 0; i < count-1; i++ {
					result = append(result, prev)
				}
			}
			prev = 0
		} else {
			// Обычный символ
			result = append(result, r)
			prev = r
		}
	}

	if isEscaped {
		return "", fmt.Errorf("некорректная строка: заканчивается символом экранирования")
	}

	return string(result), nil
}
