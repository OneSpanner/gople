package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Exer10
//
//	@Description: 练习1.10 需要入参为 “https://gopl.io https://goland.org https://godoc.org”
func Exer10() {
	fmt.Println("\n=====================Exer10 Start=====================")
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

	fmt.Println("=====================Exer10 End=====================")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
