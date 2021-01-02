package main

import (
	"bufio"
	"log"
	"os"
)

// NewLoader returns a pointer to a new instance of a TxtLoader
func NewLoader() *TxtLoader {
	loader := new(TxtLoader)
	return loader
}

func (t *TxtLoader) load(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
