package main

import (
	"fmt"
	"os"
	"github.com/benotsman-youssuf/grawl/internal/fetcher"
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

	fmt.Printf("Status: %d | Size: %d bytes\n", result.StatusCode, result.ByteCount)
}