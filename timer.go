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

// The Elapsed function prints output displaying the elapsed time since start
// in HH:MM:SS format.
func Elapsed() {
	// https://gobyexample.com/tickers
	ticker := time.NewTicker(time.Second)
	start := time.Now()
	go func() {
		// https://gobyexample.com/range-over-channels
		fmt.Printf("\r00:00:00 ")
		for range ticker.C {
			d := time.Since(start)
			printDuration(d)
		}
	}()
}

// The Countdown function prints output displaying the remaining time
// relative to the total, in HH:MM:SS format.
func Countdown(d time.Duration) {
	ticker := time.NewTicker(time.Second)
	start := time.Now()
	go func() {
		printDuration(d)
		for range ticker.C {
			remaining := d - time.Since(start)
			if remaining >= 0.0 {
				printDuration(remaining)
			} else {
				fmt.Println()
				os.Exit(0)
			}
		}
	}()
}

func exitOnNewline() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	// TODO
	//alarm := flag.String("a", "", "alarm filename")
	countdown := flag.Duration("c", time.Second*0, "countdown (duration)")
	flag.Parse()

	if *countdown >= time.Second {
		Countdown(*countdown)
	} else {
		Elapsed()
	}

	exitOnNewline()
}
