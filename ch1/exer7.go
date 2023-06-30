// Package ch1
// @Description: 第一章练习
package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Exer7
//
//	@Description: 练习1.7
func Exer7() {
	fmt.Println("\n=====================Exer7 Start=====================")

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}

	fmt.Println("=====================Exer7 End=====================")
}
