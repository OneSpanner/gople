package ch1

import (
	"fmt"
	"os"
	"time"
)

// Exer11
//
//	@Description: 练习1.11 需要入参为 “https://gopl.io https://goland.org https://godoc.org”
func Exer11() {
	fmt.Println("\n=====================Exer11 Start=====================")
	start := time.Now()
	ch := make(chan string)
	counter := 0
	for _, url := range os.Args[1:] {
		for i := 0; i < 2; i++ {
			go fetch(url, ch) // start a goroutine
			counter++
		}
	}

	for i := 0; i < counter; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	fmt.Println("=====================Exer11 End=====================")
}
