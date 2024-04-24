package Mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func TestCountDown(t *testing.T) {
	// 	buffer := &bytes.Buffer{}
	// 	spySleeper := &SpySleeper{}
	//
	// 	Countdown(buffer, spySleeper)
	//
	// 	got := buffer.String()
	// 	want := `2
	// 1
	// Go!`
	//
	// 	if got != want {
	// 		t.Errorf("Got %s, want %s", got, want)
	// 	}
	//
	// 	if spySleeper.Calls != 2 {
	// 		t.Errorf("not enough sleep calls, expected 3, got %d", spySleeper.Calls)
	// 	}
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}
