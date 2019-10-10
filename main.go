package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Heyyo...I'm still slow")
		time.Sleep(10 * time.Minute)
	}
}
