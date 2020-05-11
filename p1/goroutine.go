package main

import (
	"fmt"
	"time"
)

func main() {
	go printH()
	go printO()
	fmt.Scanln()
}

func printH() {
	for i := 1; true; i++ {
		fmt.Println(i, "H")
		time.Sleep(time.Millisecond * 200)
		fmt.Println(i, "H")
		time.Sleep(time.Millisecond * 200)
	}
}

func printO() {
	for i := 1; true; i++ {
		fmt.Println(i, "O")
		time.Sleep(time.Millisecond * 400)
	}
}
