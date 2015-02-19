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
func Countdown(d time.Duration) {
	ticker := time.NewTicker(time.Second)
	start := time.Now()
	go func() {
		printDuration(d)
		for range ticker.C {
			remaining := d - time.Since(start) + time.Second
			if remaining >= 0.0 {
				printDuration(remaining)
			} else {
				fmt.Println()
				os.Exit(0)
			}
		}
	}()
}

// The Elapsed function prints elapsed time as HH:MM:SS.
func Elapsed() {
	ticker := time.NewTicker(time.Second)
	start := time.Now()
	go func() {
		fmt.Printf("\r00:00:00 ")
		for range ticker.C {
			printDuration(time.Since(start))
		}
	}()
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

	var input string
	fmt.Scanln(&input)
}
