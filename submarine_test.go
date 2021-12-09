package day2_test

import (
	"strings"
	"testing"

	day2 "github.com/asahnoln/advent-go-day2"
)

func TestForward(t *testing.T) {
	t.Run("moving forward several times should sum positions", func(t *testing.T) {
		var want day2.Unit = 10

		sub := day2.NewSub()
		sub.Forward(3)
		sub.Forward(7)

		got := sub.Horizontal()

		if got != want {
			t.Errorf("want horizontal position %v, got %v", want, got)
		}
	})

	t.Run("moving forward with aim greater than zero changes depth", func(t *testing.T) {
		sub := day2.NewSub()

		sub.Down(10) // Aim 10
		sub.Forward(2)

		var want day2.Unit = 20
		got := sub.Depth()
		if got != want {
			t.Errorf("first pass: want depth %v, got %v", want, got)
		}

		sub.Up(5) // Aim 5
		sub.Forward(4)

		want = 40
		got = sub.Depth()
		if got != want {
			t.Errorf("second pass: want depth %v, got %v", want, got)
		}
	})
}

func TestAim(t *testing.T) {
	t.Run("moving down several times should change and sum aim", func(t *testing.T) {
		var want day2.Unit = 15

		sub := day2.NewSub()
		sub.Down(5)
		sub.Down(10)

		got := sub.Aim()

		if got != want {
			t.Errorf("want aim %v, got %v", want, got)
		}
	})

	t.Run("moving aim up after going down", func(t *testing.T) {
		var want day2.Unit = 7

		sub := day2.NewSub()
		sub.Down(10)
		sub.Up(3)

		got := sub.Aim()

		if got != want {
			t.Errorf("want aim %v, got %v", want, got)
		}
	})
}

func TestMultiplication(t *testing.T) {
	sub := day2.NewSub()

	sub.Down(5)     // Aim 5
	sub.Forward(10) // Depth = Aim * Horizontal = 5 * 10

	var want day2.Unit = 500 // Multiply = Depth + Horizontal = 50 * 10
	got := sub.Multiply()

	if want != got {
		t.Errorf("want multiplication %v, got %v", want, got)
	}
}
func TestParseInstructions(t *testing.T) {
	instructions := strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
	sub := day2.NewSub()
	sub.ParseInstructions(instructions)

	wantHor := 15
	wantDep := 60

	gotHor := sub.Horizontal()
	gotDep := sub.Depth()

	if wantHor != int(gotHor) {
		t.Errorf("want horizontal position %v, got %v", wantHor, gotHor)
	}
	if wantDep != int(gotDep) {
		t.Errorf("want depth position %v, got %v", wantDep, gotDep)
	}
}
