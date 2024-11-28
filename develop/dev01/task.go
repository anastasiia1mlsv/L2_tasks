package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// Получаем время с сервера ntp
	timeNtp, err := ntp.Time("pool.ntp.org")

	// Выводим ошибки если они есть
	if err != nil {
		b, _ := fmt.Fprintf(os.Stderr, "Error at getting time: %v\n", err)
		fmt.Printf("%d bytes written", b)
		os.Exit(1)
	}

	// Выводим время если err = nil
	fmt.Printf("Current time: %s", timeNtp.Format(time.RFC3339))
}
