package main

import "fmt"

func main() {
	u := make(chan int, 1)
	u <- 3
	go func() {
		for t := range u {
			fmt.Println(t)
		}
	}()

}
