package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	co := make(chan string)

	go printH(ch)
	go printO(co)

	for {
		select {
		case <-ch:
			fmt.Println(<-co)
		case <-co:
			fmt.Println(<-ch)
			fmt.Println(<-ch)
		}
	}
}

func printH(ch chan string) {
	for {
		ch <- "H"
		time.Sleep(time.Millisecond * 200)
	}
}

func printO(co chan string) {
	for {
		co <- "O"
		time.Sleep(time.Millisecond * 200)
	}
}
