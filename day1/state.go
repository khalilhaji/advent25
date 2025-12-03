package main

import (
	"math"
)

type state struct {
	counter int
}

func (s state) Execute(instr instruction) (state, int) {
	var newCounter int
	passedZero := 0
	if instr.dir == DIRECTION_LEFT {
		newCounter = s.counter - instr.num

		if newCounter <= 0 {
			if s.counter == 0 {
				passedZero += instr.num / 100
			} else {
				passedZero = ((instr.num - s.counter) / 100) + 1
			}
		}
	} else {
		newCounter = s.counter + instr.num
		passedZero = (instr.num + s.counter) / 100
	}

	moddedCounter := math.Mod(float64(newCounter), 100)
	if moddedCounter < 0 {
		moddedCounter += 100
	}

	return state{
		counter: int(moddedCounter),
	}, passedZero
}

func CountZeroEndStates(instructions []instruction) int {
	total := 0
	s := state{50}
	for _, instr := range instructions {
		s, _ = s.Execute(instr)
		if s.counter  == 0 {
			total++
		}
	}

	return total
}

func CountZeroPasses(instructions []instruction) int {
	total := 0
	s := state{50}
	for _, instr := range instructions {
		newState, passedZero := s.Execute(instr)
		s = newState
		total += passedZero
	}

	return total
}
