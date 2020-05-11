package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	go func() {
		for i := 1; true; i++ {
			wg.Add(3)
			printH(i)
			printH(i)
			printO(i)
			wg.Done()
		}
	}()

	fmt.Scanln()
}

func printH(i int) {
	fmt.Println(i, "H")
	time.Sleep(time.Millisecond * 200)
}

func printO(i int) {
	fmt.Println(i, "O")
	time.Sleep(time.Millisecond * 200)
}
