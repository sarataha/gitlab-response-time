package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	interval := flag.Uint("interval", 5, "Request interval in seconds")
	duration := flag.Uint("duration", 5, "Running duration in minutes")
	url := flag.String("url", "https://gitlab.com", "URL under examination")
	flag.Parse()

	go func() {
		time.Sleep(time.Duration(*duration) * time.Minute)
		os.Exit(0)
	}()

	i := 0
	for {
		go func() {
			n := i
			start := time.Now()
			_, err := http.Get(*url)

			if err != nil {
				fmt.Println(n, "- Error connecting to host")
			} else {
				elapsed := time.Since(start)
				fmt.Println(n, "-", elapsed)
			}
		}()
		time.Sleep(time.Duration(*interval) * time.Second)
		i++
	}
}
