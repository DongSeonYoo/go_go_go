package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", spoon, fork, "수저", "포크")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName string, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으러 갑니다\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득 \n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득 \n", name, secondName)

		fmt.Printf("%s 밥을 먹스빈다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))

		second.Unlock()
		first.Unlock()
	}

	wg.Done()
}
