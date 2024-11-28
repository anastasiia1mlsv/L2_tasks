### Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```
### Ответ: error

- Функция test() возвращает nil типа *customError.
- В main(), переменная err имеет тип error и получает значение test().
- Хотя test() возвращает nil, переменная err не равна nil, потому что она содержит интерфейс error с типом *customError и значением nil.
- При сравнении err != nil, результат будет true, потому что интерфейс не nil (у него есть тип).
- Поэтому программа выводит "error" и завершает выполнение.
