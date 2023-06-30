package ch1

import (
	"bufio"
	"fmt"
	"os"
)

// Exer4
//
//	@Description: 练习1.4
func Exer4() {
	fmt.Println("\n=====================Exer4 Start=====================")

	counts := make(map[string]int)
	files := "./ch1/exer4ex.txt"
	f, err := os.Open(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exer4: %v\n", err)
	}
	countLines(f, counts)
	err = f.Close()
	if err != nil {
		return
	}

	fmt.Println("=====================Exer4 End=====================")
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("文件%v中%v行出现次数为%v\n", f.Name(), input.Text(), counts[input.Text()])
		}
	}
}
