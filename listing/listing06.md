### Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```
### Ответ: [3 2 3]

- Начальное состояние слайса:
s = []string{"1", "2", "3"}.
- Передача слайса в modifySlice:
В Go слайсы передаются по значению, но их внутреннее устройство включает ссылку на подлежащий массив, длину (len) и ёмкость (cap).
- Вызов modifySlice:
i[0] = "3": Меняет значение первого элемента исходного массива (слайс s ссылается на тот же массив). Теперь массив ["3", "2", "3"].
i = append(i, "4"): Если ёмкость подлежащего массива позволяет, изменяется тот же массив. Если ёмкости недостаточно, создаётся новый массив, куда копируются данные. В данном случае ёмкость изначального массива достаточна, и новый элемент добавляется, но изменения касаются только локального слайса i.
i[1] = "5": Меняет второй элемент локального слайса i. Но так как i ссылается на новый массив, это не затрагивает исходный слайс s.
i = append(i, "6"): Снова добавляется новый элемент в локальный слайс i.
- Возврат к main:
Изменения, затронувшие исходный массив (i[0] = "3"), сохранились в s.
Остальные изменения (в результате append) не затронули s, так как локальный слайс i стал ссылаться на новый массив.





Внутреннее устройство слайсов:

Слайс состоит из указателя на подлежащий массив, 
длины (len), 
указывающей на количество доступных элементов и
емкости (cap), указывающей 
на максимальное количество элементов 
до выделения новой памяти.


Передача слайса в функцию:

Передаётся копия структуры слайса, 
но копия ссылается на тот же подлежащий массив.
Изменение элементов через индекс влияет 
на исходный массив.
Если слайс изменяет размер 
(например, через append), 
он может начать ссылаться на новый массив, 
и дальнейшие изменения не затрагивают исходный массив.
