package main

import (
	"fmt"
	"os"
	"time"
)

const usage string = `A simple CLI stopwatch.
Usage: 4d [DURATION]

4d		display elapsed time
4d 15m		countdown 15 minutes
4d 3h2m1s	countdown 3 hours, 2 minutes, 1 second

Ctrl-C exits.`

// printDuration prints a given duration as HH:MM:SS.
func printDuration(d time.Duration) {
	fmt.Printf("\r%02d:%02d:%02d ",
		int(d.Hours())%60,
		int(d.Minutes())%60,
		int(d.Seconds())%60,
	)
}

// Countdown prints time remaining relative to a given total.
func Countdown(ticker *time.Ticker, d time.Duration) {
	start := time.Now()
	for range ticker.C {
		remaining := d - time.Since(start) + time.Second
		if remaining >= 0.0 {
			printDuration(remaining)
		} else {
			return
		}
	}
}

// Elapsed prints the duration since the provided start time.
func Elapsed(ticker *time.Ticker, start time.Time) {
	for range ticker.C {
		printDuration(time.Since(start))
	}
}

func main() {
	var countdown time.Duration

	// parse duration if provided
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "-help", "--help", "help":
			fmt.Printf("\n%s\n\n", usage)
			os.Exit(0)
		}

		d, err := time.ParseDuration(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		countdown = d
	}

	// - set tick lower than than the smallest unit of interest (1s) to
	//   mitigate output "skips" due to missed ticks
	// - related issue: https://github.com/golang/go/issues/3516
	ticker := time.NewTicker(time.Second / 2)

	defer fmt.Println()
	if countdown >= time.Second {
		Countdown(ticker, countdown)
	} else {
		Elapsed(ticker, time.Now())
	}
}
