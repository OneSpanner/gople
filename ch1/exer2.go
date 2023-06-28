// Package ch1
// @Description: 第一章练习
package ch1

import (
	"fmt"
	"os"
)

import (
	"strconv"
)

// Exer2
//
//	@Description: 练习1.2
func Exer2() {
	fmt.Println("\n=====================Exer2 Start=====================")

	for i, arg := range os.Args {
		fmt.Println("arg " + strconv.Itoa(i) + ": " + arg)
	}

	fmt.Println("=====================Exer2 End=====================")
}
