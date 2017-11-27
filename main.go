package main

import (
	"fmt"
	"strings"
	"time"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)

	for _, char := range strings.Fields(s) {
		count[char]++
	}

	return count
}

func main() {
	start := time.Now()
	pharses := []string{
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
		"Hello my name is Jeff",
	}

	//fast(pharses)
	slow(pharses)

	fmt.Println(time.Since(start))

}

func slow(pharses []string) {
	count := make(map[string]int)

	for _, pharse := range pharses {
		for k, v := range WordCount(pharse) {
			count[k] += v
		}
	}

	fmt.Println(count)
}

func fast(pharses []string) {
	count := make(map[string]int)
	ch := make(chan map[string]int)

	for _, pharse := range pharses {
		go func() {
			ch <- WordCount(pharse)
		}()
	}

	for i := 0; i < len(pharses); i++ {
		res := <-ch
		for k, v := range res {
			count[k] += v
		}
	}

	fmt.Println(count)
}
