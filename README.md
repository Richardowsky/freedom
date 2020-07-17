# Find the Freedom

## Prerequisites:
You need to install - [go 1.13](https://golang.org/dl/)

## How to test:
1. Clone this repo: `git clone https://github.com/Richardowsky/freedom.git`
2. `make test`

## How to use:
```go
package example

import freedom "freedom/src"

func example()  {
	
var n = 4
var k = 4
var array = []int{1, 1, 2, 2}
	
	freedom.Solution(n, k, array)
}

```

Алгоритм:
- сортируем масив
- проверяем кратно ли значение для k и исключаем все кратные элементы
- проверяем следущие элементы по возрастанию, делением на k
- распределяем сначала в первую половину масива четные, а во вторую нечетные значения 

Алгорим сортировки  = n\log n
Алгоритм решения  = n