package countdown

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *SpyCountdownOperations) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countDownStart = 3

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}
