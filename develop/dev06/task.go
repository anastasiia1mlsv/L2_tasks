package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определяем флаги командной строки
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель (по умолчанию TAB)")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Проверяем, что флаг -f указан
	if *fields == "" {
		_, err := fmt.Fprintln(os.Stderr, "cut: необходимо указать поля с помощью -f")
		if err != nil {
			return
		}
		os.Exit(1)
	}

	// Парсим номера колонок из флага -f
	fieldList, err := parseFields(*fields)
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, "cut:", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	// Читаем строки из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Если установлен флаг -s, пропускаем строки без разделителя
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разбиваем строку по разделителю
		columns := strings.Split(line, *delimiter)

		// Извлекаем указанные колонки
		var output []string
		for _, idx := range fieldList {
			// Проверяем, что индекс находится в пределах доступных колонок
			if idx-1 < len(columns) && idx-1 >= 0 {
				output = append(output, columns[idx-1])
			}
		}

		// Печатаем выбранные колонки
		fmt.Println(strings.Join(output, *delimiter))
	}

	// Обрабатываем ошибки чтения
	if err := scanner.Err(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "cut: ошибка при чтении:", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

// parseFields парсит строку с номерами полей и возвращает список индексов
func parseFields(fields string) ([]int, error) {
	var fieldList []int
	parts := strings.Split(fields, ",") // Разделяем строку по запятым
	for _, part := range parts {
		// Поддержка диапазонов, например, 1-3
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("некорректный диапазон: %s", part)
			}
			start, err := strconv.Atoi(rangeParts[0])
			if err != nil || start <= 0 {
				return nil, fmt.Errorf("некорректное начальное значение диапазона: %s", part)
			}
			end, err := strconv.Atoi(rangeParts[1])
			if err != nil || end < start {
				return nil, fmt.Errorf("некорректное конечное значение диапазона: %s", part)
			}
			for i := start; i <= end; i++ {
				fieldList = append(fieldList, i)
			}
		} else {
			// Одиночный номер поля
			idx, err := strconv.Atoi(part)
			if err != nil || idx <= 0 {
				return nil, fmt.Errorf("некорректный номер поля: %s", part)
			}
			fieldList = append(fieldList, idx)
		}
	}
	return fieldList, nil
}
