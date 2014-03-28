package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"io"
	"os"
)

var data = flag.String("data", "", "")

var graph *map[word]*list.List

func parse(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
}

func main() {
	flag.Parse()

	dataFile, _ := os.Open(*data)

	fmt.Println("hej")
}
