package main

import (
	"fmt"
	"strconv"
	"time"
)

func chan1Run(ch chan string) {
	for i := 0; i < 19; i++ {
		ch <- "i=" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
}
func chan2Run(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go chan1Run(ch1)
		go chan2Run(ch2)
	}
	for {
		// 1次select 只会出一次结果
		select {
		case str, ch1Check := <-ch1:
			if !ch1Check {
				fmt.Println("ch1Check false")
			}
			fmt.Println(str)
		case p, ch2Check := <-ch2:
			if !ch2Check {
				fmt.Println("ch2Check false")
			}
			fmt.Println(p)
		}
	}

	time.Sleep(60 * time.Second)
}
