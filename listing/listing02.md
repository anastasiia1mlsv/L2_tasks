### Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```
### Ответ: 2, 1

- В функции test() переменная x является 
именованным возвращаемым параметром. 
При выполнении return значение x фиксируется, 
затем выполняется defer, который увеличивает x на 1. 
Поэтому возвращается 2.


- В функции anotherTest() переменная x 
не является именованным возвращаемым параметром. 
Значение x возвращается как 1 до выполнения defer, 
который увеличивает x, но это не влияет на уже 
возвращённое значение. Поэтому возвращается 1.
