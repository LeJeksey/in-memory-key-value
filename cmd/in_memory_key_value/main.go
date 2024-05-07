package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/LeJeksey/in-memory-key-value/internal/compute"
	"github.com/LeJeksey/in-memory-key-value/internal/storage"
)

func main() {
	store := storage.NewStorage()

	analyzer := compute.NewAnalyzer()
	parser := compute.NewParserSM()

	computer := compute.NewComputer(store, parser, analyzer)

	// read queries from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		query := scanner.Text()

		result, err := computer.Compute(query)
		if err != nil {
			log.Printf("failed to compute query: %v", err)
			continue
		}

		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
}
