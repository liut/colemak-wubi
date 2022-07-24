package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	replColemakDH = strings.NewReplacer(
		"e", "f", "r", "p", "t", "b",
		"y", "j", "u", "l", "i", "u", "o", "y", "p", "o",
		"s", "r", "d", "s", "f", "t",
		"h", "m", "j", "n", "k", "e", "l", "i",
		"v", "d", "b", "v", "n", "k", "m", "h")
)

func main() {
	var (
		input  string
		output string
	)
	flag.StringVar(&input, "i", "", "input")
	flag.StringVar(&output, "o", "", "output")

	flag.Parse()
	if input == "" || output == "" {
		flag.Usage()
		return
	}

	if input == output {
		log.Print("Don't overwrite original file")
		flag.Usage()
		return
	}

	fin, err := os.Open(input)
	if err != nil {
		log.Printf("open fail: %s", err)
		return
	}
	defer fin.Close()
	fout, err := os.Create(output)
	if err != nil {
		log.Printf("create fail: %s", err)
		return
	}
	defer fout.Close()
	scanner := bufio.NewScanner(fin)
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(replColemakDH.Replace(line)) // Println will add back the final '\n'
		fmt.Fprintln(fout, replColemakDH.Replace(line))
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Printf("reading input err: %s", err)
	}
	log.Printf("%d lines", count)
}
