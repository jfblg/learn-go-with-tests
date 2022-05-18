package main

import (
	"mocking"
	"os"
	"time"
)

func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}

	mocking.Countdown(os.Stdout, sleeper)
}
