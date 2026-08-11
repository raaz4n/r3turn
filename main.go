package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./programname <file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	defer file.Close()

	client := &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// check if the string contains https://
		if !strings.Contains(line, "https://") {
			line = "https://" + line
		}

		resp, err := client.Get(line)
		if err != nil {
			fmt.Printf("Failed to send req to %s\n", line)
			continue
		}

		defer resp.Body.Close()

		status := resp.Status

		fmt.Printf("[%s] %s\n", status, line)
	}
}
