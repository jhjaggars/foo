package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I'm foo")
		time.Sleep(1 * time.Minute)
	}
}
