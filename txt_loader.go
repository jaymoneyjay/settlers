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

func (t *TxtLoader) load(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t.lines = append(t.lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
