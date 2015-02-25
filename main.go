package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func printDuration(d time.Duration) {
	fmt.Printf("\r%02d:%02d:%02d ",
		int(d.Hours())%60,
		int(d.Minutes())%60,
		int(d.Seconds())%60,
	)
}

// The Countdown function prints time remaining relative to a given total (as HH:MM:SS).
func Countdown(ticker *time.Ticker, d time.Duration) {
	start := time.Now()
	for range ticker.C {
		remaining := d - time.Since(start) + time.Second
		if remaining >= 0.0 {
			printDuration(remaining)
		} else {
			fmt.Println()
			os.Exit(0)
		}
	}
}

// The Elapsed function prints elapsed time as HH:MM:SS.
func Elapsed(ticker *time.Ticker) {
	start := time.Now()
	for range ticker.C {
		printDuration(time.Since(start))
	}
}

func main() {
	countdown := flag.Duration("c", time.Second*0, "countdown (duration)")
	flag.Parse()

	// - set tick lower than than the smallest unit of interest (1s) to
	// mitigate output "skips" due to missed ticks
	// - related issue: https://github.com/golang/go/issues/3516
	ticker := time.NewTicker(time.Second / 2)

	if *countdown >= time.Second {
		go Countdown(ticker, *countdown)
	} else {
		go Elapsed(ticker)
	}

	var input string
	fmt.Scanln(&input)
}
