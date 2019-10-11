package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I'm just gonna lay down...")
		time.Sleep(1 * time.Hour)
	}
}
