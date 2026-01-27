package main

import (
	"fmt"
	"sync"
	"time"
)

func computeSquare(num int, resChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	resChan <- num * num
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resChan := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		go computeSquare(number, resChan, &wg)
	}
	wg.Wait()
	close(resChan)
	for square := range resChan {
		fmt.Println(square)
	}
}
