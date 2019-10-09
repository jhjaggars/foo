package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I'm slow foo")
		time.Sleep(10 * time.Minute)
	}
}
