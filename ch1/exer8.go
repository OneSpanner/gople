// Package ch1
// @Description: 第一章练习
package ch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Exer7
//
//	@Description: 练习1.8
func Exer8() {
	fmt.Println("\n=====================Exer8 Start=====================")

	//tempArgs := []string{"", "www.baidu.com"}
	//for _, url := range tempArgs[1:] {

	for _, url := range os.Args[1:] {
		// joint prefix "http://"
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}

	fmt.Println("=====================Exer8 End=====================")
}
