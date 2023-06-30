// Package ch1
// @Description: 第一章练习
package ch1

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Exer9
//
//	@Description: 练习1.9
func Exer9() {
	fmt.Println("\n=====================Exer9 Start=====================")

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
		//b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)
		fmt.Printf("Status =%v\n", resp.Status)
	}

	fmt.Println("=====================Exer9 End=====================")
}
