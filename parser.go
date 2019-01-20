package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	Parse(3, "test", "test2")
}

func Parse(num int, urls ...string) {
	//var a []Event
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			continue
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text)
		}
	}
}
