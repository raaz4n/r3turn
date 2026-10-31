// r3turn by raaz4n
// The usage for this program is ./r3turn <file name> <request time in milliseconds>
// It is most efficient when used alongside subfinder.

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./r3turn <file> <request time in ms>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	defer file.Close()

	timeMs, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid ms value: ", err)
		return
	}

	client := &http.Client{
		Timeout: time.Duration(timeMs) * time.Millisecond,
	}

	scanner := bufio.NewScanner(file)

	sem := make(chan struct{}, 20)
	var wg sync.WaitGroup

	fmt.Println("In progress. All requests that return errors will be discarded. If no results are found, try increasing the ms.\n")

	for scanner.Scan() {
		line := scanner.Text()
		// check if the string contains https://
		if !strings.Contains(line, "https://") {
			line = "https://" + line
		}

		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := client.Get(line)
			if err != nil {
				return
			}

			defer resp.Body.Close()

			status := resp.Status

			fmt.Printf("[%s] %s\n", status, line)
		}()
	}

	wg.Wait()

	fmt.Println("\nDone!")
}
