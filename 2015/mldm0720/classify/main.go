package main

import (
	"bufio"
	"fmt"
	"github.com/guesslin/mits"
	"os"
	"strings"
)

func SegFile(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	strs := make([][]rune, 0)
	for scanner.Scan() {
		strs = append(strs, []rune(strings.ToLower(scanner.Text())))
	}
	terms := mits.SegmentSens(strs, 0.001)
	return terms
}

func main() {
	fmt.Println("---------Normal--------")
	for _, c := range SegFile("./top500_top.txt") {
		if len(c) > 3 {
			fmt.Printf("%s ", c)
		}
	}
	fmt.Println() // OMIT
	fmt.Println("-------Malicious-------")
	for _, m := range SegFile("./malicious_top.txt") {
		if len(m) > 3 {
			fmt.Printf("%s ", m)
		}
	}
	fmt.Println() // OMIT
}
