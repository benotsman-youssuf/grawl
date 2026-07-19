package main

import (
	"fmt"
	"os"
	"github.com/benotsman-youssuf/grawl/internal/fetcher"
	"github.com/benotsman-youssuf/grawl/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing URL")
		os.Exit(1)
	}

	targetUrl := os.Args[1]

	result, err := fetcher.FetchUrl(targetUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Execution failed: %v\n", err)
		os.Exit(1)
	}

	for _, link := range parser.ParseUrl(result.ByteCount) {
		fmt.Println(link)
	}
}