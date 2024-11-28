package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// GrepOptions содержит опции для утилиты grep
type GrepOptions struct {
	After      int    // -A N: печатать +N строк после совпадения
	Before     int    // -B N: печатать +N строк до совпадения
	Context    int    // -C N: печатать ±N строк вокруг совпадения
	Count      bool   // -c: печатать количество строк
	IgnoreCase bool   // -i: игнорировать регистр
	Invert     bool   // -v: исключать совпадения
	Fixed      bool   // -F: точное совпадение со строкой
	LineNum    bool   // -n: печатать номер строки
	Pattern    string // Искомый шаблон
	FilePath   string // Путь к файлу
}

func main() {
	// Создаем буфер для чтения данных из консоли
	reader := bufio.NewReader(os.Stdin)

	// Выводим заголовок
	fmt.Println("Grep Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("command> ")

		// Считываем строку из буфера
		buf, _ := reader.ReadString('\n')
		buf = strings.TrimSpace(buf)

		// Разбиваем введенную строку на аргументы
		args := parseArgs(buf)

		// Проверяем, что введена команда и необходимое количество аргументов
		if len(args) < 3 {
			fmt.Println("Недостаточно аргументов. Использование: grep [OPTIONS] PATTERN FILE")
			continue
		}

		command := args[0]

		// Проверяем команду на 'exit' для выхода из программы
		if command == "exit" {
			os.Exit(0)
		}

		// Проверяем, что команда 'grep'
		if command != "grep" {
			fmt.Println("Неизвестная команда:", command)
			continue
		}

		// Парсим опции и получаем структуру GrepOptions
		opts, err := parseGrepOptions(args[1:])
		if err != nil {
			fmt.Println("Ошибка при разборе опций:", err)
			continue
		}

		// Выполняем поиск
		err = grepFile(opts)
		if err != nil {
			fmt.Println("Ошибка при выполнении grep:", err)
			continue
		}
	}
}

// parseArgs разбивает строку на аргументы, учитывая кавычки
func parseArgs(input string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false

	for _, r := range input {
		switch r {
		case ' ':
			if inQuotes {
				current.WriteRune(r)
			} else {
				if current.Len() > 0 {
					args = append(args, current.String())
					current.Reset()
				}
			}
		case '"':
			inQuotes = !inQuotes
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

// parseGrepOptions парсит аргументы и возвращает структуру GrepOptions
func parseGrepOptions(args []string) (*GrepOptions, error) {
	opts := &GrepOptions{}
	i := 0
	for i < len(args)-2 { // Последние два аргумента - PATTERN и FILE
		arg := args[i]
		switch arg {
		case "-A":
			i++
			if i >= len(args)-2 {
				return nil, fmt.Errorf("флаг -A требует аргумент")
			}
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return nil, fmt.Errorf("не удалось преобразовать %q в число: %v", args[i], err)
			}
			opts.After = n
		case "-B":
			i++
			if i >= len(args)-2 {
				return nil, fmt.Errorf("флаг -B требует аргумент")
			}
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return nil, fmt.Errorf("не удалось преобразовать %q в число: %v", args[i], err)
			}
			opts.Before = n
		case "-C":
			i++
			if i >= len(args)-2 {
				return nil, fmt.Errorf("флаг -C требует аргумент")
			}
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return nil, fmt.Errorf("не удалось преобразовать %q в число: %v", args[i], err)
			}
			opts.Context = n
		case "-c":
			opts.Count = true
		case "-i":
			opts.IgnoreCase = true
		case "-v":
			opts.Invert = true
		case "-F":
			opts.Fixed = true
		case "-n":
			opts.LineNum = true
		default:
			return nil, fmt.Errorf("неизвестный флаг: %s", arg)
		}
		i++
	}

	// Последние два аргумента - шаблон и имя файла
	opts.Pattern = args[len(args)-2]
	opts.FilePath = args[len(args)-1]

	return opts, nil
}

// grepFile выполняет поиск в файле по заданным опциям
func grepFile(opts *GrepOptions) error {
	// Открываем файл для чтения
	file, err := os.Open(opts.FilePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Читаем строки из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	// Создаем регулярное выражение для поиска
	pattern := opts.Pattern
	if opts.Fixed {
		// Экранируем специальные символы для точного совпадения
		pattern = regexp.QuoteMeta(pattern)
	}
	flags := ""
	if opts.IgnoreCase {
		flags = "(?i)"
	}
	regex, err := regexp.Compile(flags + pattern)
	if err != nil {
		return fmt.Errorf("не удалось скомпилировать регулярное выражение: %v", err)
	}

	// Массив для хранения номеров строк, которые совпали
	var matchedLines []int

	// Ищем совпадения
	for i, line := range lines {
		matched := regex.MatchString(line)
		if opts.Invert {
			matched = !matched
		}
		if matched {
			matchedLines = append(matchedLines, i)
		}
	}

	// Если установлен флаг -c, выводим количество совпадений
	if opts.Count {
		fmt.Println(len(matchedLines))
		return nil
	}

	// Создаем множество для хранения номеров строк, которые нужно вывести
	linesToPrint := make(map[int]struct{})

	// Добавляем строки с учетом флагов -A, -B, -C
	for _, lineNum := range matchedLines {
		start := lineNum
		end := lineNum

		// Обрабатываем флаг -C
		if opts.Context > 0 {
			start = lineNum - opts.Context
			end = lineNum + opts.Context
		} else {
			// Обрабатываем флаги -A и -B
			if opts.Before > 0 {
				start = lineNum - opts.Before
			}
			if opts.After > 0 {
				end = lineNum + opts.After
			}
		}

		// Ограничиваем диапазон
		if start < 0 {
			start = 0
		}
		if end >= len(lines) {
			end = len(lines) - 1
		}

		// Добавляем номера строк в множество
		for i := start; i <= end; i++ {
			linesToPrint[i] = struct{}{}
		}
	}

	// Выводим строки
	for i, line := range lines {
		if _, ok := linesToPrint[i]; ok {
			if opts.LineNum {
				fmt.Printf("%d:%s\n", i+1, line)
			} else {
				fmt.Println(line)
			}
		}
	}

	return nil
}
