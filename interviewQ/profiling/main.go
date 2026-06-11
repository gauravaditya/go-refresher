package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func buildString(n int) string {
	var s string

	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%d,", i)
	}

	return s
}

func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	results := make([]string, 0, 1000)

	for i := 0; i < 1000; i++ {
		results = append(results, buildString(1000))
	}
	pprof.WriteHeapProfile(f)

	fmt.Println("done")
}
