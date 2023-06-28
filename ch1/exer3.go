package ch1

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Exer3
//
//	@Description: 练习1.3
func Exer3() {
	fmt.Println("\n=====================Exer3 Start=====================")

	var args []string
	for i := 0; i < 100000; i++ {
		args = append(args, strconv.Itoa(i))
	}
	//fmt.Printf("args := %v", args)

	start1 := time.Now()
	connectStrWithJoin(args...)
	time1 := time.Since(start1)
	fmt.Printf("使用Join方法连接%d个元素花费时长%s\n", len(args), time1)

	start2 := time.Now()
	connectStrWithoutJoint(args...)
	time2 := time.Since(start2)
	fmt.Printf("不使用Join方法连接%d个元素花费时长%s\n", len(args), time2)

	fmt.Println("=====================Exer3 End=====================")
}

// connectStrWithJoin
//
//	@Description: 使用strings.Join方法
//	@param args...
//	@return string
func connectStrWithJoin(args ...string) string {
	return strings.Join(args, " ")
}

// connectStrWithoutJoint
//
//	@Description: 不使用strings.Join方法
//	@param args...
//	@return string
func connectStrWithoutJoint(args ...string) string {
	if len(args) <= 0 {
		return ""
	}
	ret := args[0]
	for _, arg := range args[1:] {
		ret += arg
	}
	return ret
}
