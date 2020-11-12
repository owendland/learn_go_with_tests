package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration: duration, sleep: sleep}
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) error {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()

		_, err := fmt.Fprintln(out, i)
		if err != nil {
			return err
		}
	}

	sleeper.Sleep()

	_, err := fmt.Fprint(out, finalWord)
	if err != nil {
		return err
	}

	return nil
}
