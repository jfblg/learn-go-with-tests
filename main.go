package main

import (
	"mocking"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}

	mocking.Countdown(os.Stdout, sleeper)
}
