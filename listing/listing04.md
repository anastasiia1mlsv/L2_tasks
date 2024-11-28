### Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```
### Ответ:

- Программа выведет числа от 0 до 9, 
а затем возникнет ошибка deadlock, 
потому что канал не закрыт 
и цикл for n := range ch бесконечно ожидает новые данные.