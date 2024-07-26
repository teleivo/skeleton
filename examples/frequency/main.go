// Frequency is a Go like version of https://algs4.cs.princeton.edu/31elementary/FrequencyCounter.java.html
// Try it with a data sample from https://introcs.cs.princeton.edu/java/data/ like dickens.txt
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/teleivo/skeleton/order"
)

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Printf("exits due to error: %v\n", err)
		os.Exit(1)
	}
}

func run(w io.Writer) error {
	fileName := flag.String("file", "", "path to input file")
	minChars := flag.Uint("minChars", 4, "a words minimum number of characters for it to be counted")
	flag.Parse()

	if *fileName == "" {
		return errors.New("-file must be specified")
	}

	f, err := os.Open(*fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	var total, totalMin int
	st := order.Table[string, int]{}
	for sc.Scan() {
		word := sc.Text()

		total++
		if utf8.RuneCount([]byte(word)) >= int(*minChars) {
			continue
		}
		totalMin++

		count, ok := st.Get(word)
		if !ok {
			st.Put(word, 1)
		} else {
			count++
			st.Put(word, count)
		}
	}

	if sc.Err() != nil {
		return fmt.Errorf("failed to scan words: %v", err)
	}

	// TODO continue once I have implemented an iterator
	fmt.Fprintf(w, "%q contains %d words with %d words with at least %d characters", *fileName, total, totalMin, *minChars)

	return nil
}
