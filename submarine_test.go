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
}

func TestDepth(t *testing.T) {
	t.Run("moving deep several times should sum positions", func(t *testing.T) {
		var want day2.Unit = 15

		sub := day2.NewSub()
		sub.Down(5)
		sub.Down(10)

		got := sub.Depth()

		if got != want {
			t.Errorf("want horizontal position %v, got %v", want, got)
		}
	})

	t.Run("moving up after going down", func(t *testing.T) {
		var want day2.Unit = 7

		sub := day2.NewSub()
		sub.Down(10)
		sub.Up(3)

		got := sub.Depth()

		if got != want {
			t.Errorf("want horizontal position %v, got %v", want, got)
		}
	})
}

func TestMultiplication(t *testing.T) {
	sub := day2.NewSub()

	sub.Forward(15)
	sub.Down(10)

	var want day2.Unit = 150
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
	wantDep := 10

	gotHor := sub.Horizontal()
	gotDep := sub.Depth()

	if wantHor != int(gotHor) {
		t.Errorf("want horizontal position %v, got %v", wantHor, gotHor)
	}
	if wantDep != int(gotDep) {
		t.Errorf("want depth position %v, got %v", wantDep, gotDep)
	}
}
