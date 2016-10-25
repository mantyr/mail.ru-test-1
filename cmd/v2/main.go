package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	s := flag.String("search", "Go", "Search phrase")
	k := flag.Int64("k", 5, "k=5 simultaneously")
	flag.Parse()

	total := Scaner(*s, *k, os.Stdin)
	fmt.Println("Total: ", total)
}

func Scaner(sep string, k int64, file *os.File) (total int64) {
	counter := NewCounter(k)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counter.Run(scanner.Text(), sep)
	}
	counter.Wait()
	return counter.GetTotal()
}
