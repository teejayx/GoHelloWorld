package main

import (
	"bytes"
	"testing"

	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)




func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
    spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
	2
	1
	G0!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
