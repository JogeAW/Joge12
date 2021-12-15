package main

import (
	"fmt"
	"Joge12/numbers"
	"sync"
)

const (
	limit   = 100000
	rutines = 10
)

func Rutines(id int, c *counter.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		if ok := c.Add(1); !ok {
			fmt.Println("Rutine:", id, "Iterations:", i)
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	// Создаем счетчик с заданным лимитом
	c := counter.NewCounter(limit)

	// Запускаем требуемое кол-во горутин-воркеров и ожидаем их завершения
	wg.Add(rutines)
	for id := 0; id < rutines; id++ {
		go Rutines(id, c, &wg)
	}
	wg.Wait()

	// Печатаем значение счетчика
	fmt.Println("Counter:", c.Value())
}
