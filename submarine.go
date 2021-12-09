package day2

import (
	"fmt"
	"io"
)

type Unit int

type Sub struct {
	horizontal Unit
	depth      Unit
	aim        Unit
}

func NewSub() *Sub {
	return &Sub{}
}

func (s *Sub) Horizontal() Unit {
	return s.horizontal
}

func (s *Sub) Depth() Unit {
	return s.depth
}

func (s *Sub) Aim() Unit {
	return s.aim
}

func (s *Sub) Forward(x Unit) {
	s.horizontal += x
	s.depth += x * s.aim
}

func (s *Sub) Down(x Unit) {
	s.aim += x
}

func (s *Sub) Up(x Unit) {
	s.Down(-x)
}

func (s *Sub) Multiply() Unit {
	return s.Horizontal() * s.Depth()
}

func (s *Sub) ParseInstructions(r io.Reader) {
	var cmd string
	var x Unit

	for {
		_, err := fmt.Fscanln(r, &cmd, &x)
		if err != nil {
			break
		}

		switch cmd {
		case "forward":
			s.Forward(x)
		case "down":
			s.Down(x)
		case "up":
			s.Up(x)
		}
	}

}
