package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var isNumbered bool

func init() {
	flag.BoolVar(&isNumbered, "n", false, "Number the output lines, starting at 1.")

}

func main() {
	flag.Parse()

	n := 1
	for _, fileName := range flag.Args() {
		if isNumbered {
			n = readFileWithNumber(fileName, n)
		} else {
			readFile(fileName)
		}
	}
}

func readFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイルのオープンに失敗しました: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました: ", err)
	}
}

func readFileWithNumber(fileName string, n int) int {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイルのオープンに失敗しました: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, strconv.Itoa(n)+": "+scanner.Text())
		n++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました: ", err)
	}
	return n
}
