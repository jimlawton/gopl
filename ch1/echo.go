package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	echo1()
	time1 := time.Since(start).Nanoseconds()
	start = time.Now()
	echo2()
	time2 := time.Since(start).Nanoseconds()
	start = time.Now()
	echo3()
	time3 := time.Since(start).Nanoseconds()

	fmt.Printf("Echo 1: %d nanoseconds\n", time1)
	fmt.Printf("Echo 2: %d nanoseconds\n", time2)
	fmt.Printf("Echo 3: %d nanoseconds\n", time3)
}
