package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func createElements(elementsCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	slice := make([]int, 10)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		slice[i] = generator.Intn(100)
	}

	for _, element := range slice {
		elementsCh <- element
	}
	close(elementsCh)
}

func toSquare(elementsCh chan int, mainCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for element := range elementsCh {
		result := element * element
		mainCh <- result
	}

	close(mainCh)
}


func main() {
	elementsCh := make(chan int)
	mainCh := make(chan int)
	slice := []int{}

	var wg sync.WaitGroup
	// 2 горутины
	wg.Add(2)
	

	go createElements(elementsCh, &wg)
	go toSquare(elementsCh, mainCh, &wg)

	for element := range mainCh {
		slice = append(slice, element)
	}
	wg.Wait()

	fmt.Println("Конец программы", slice)
}