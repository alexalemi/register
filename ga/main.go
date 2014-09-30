package main

import (
	"fmt"
	/* "time" */)

type Command int
type Register int
type Label int

const MaxReg = 16
const MaxLen = 128
const TimeOut = 1

// The commands are ints
const (
	END Command = iota
	INC
	DEC
)

// A single instruction
type Instruction struct {
	inst Command
	reg  Register
	next Label
	jump Label
}

// A whole program
type Program [MaxLen]Instruction
type World [MaxReg]Register

type Expectation struct {
	in  World
	out World
}
type Recipe []Expectation

type ScoreCard struct {
	Len        int
	Correct    int
	AllCorrect bool
	Steps      int
	TimeOut    bool
}

func (s ScoreCard) Score() float64 {
	score := 0.0
	score += float64(s.Len)
	score += float64(s.Steps)
	score -= 100 * float64(s.Correct)
	if s.AllCorrect {
		score -= 10000.0
	}
	if s.TimeOut {
		score += 10000.0
	}
	return score
}

type Organism struct {
	Prog      Program
	ScoreCard ScoreCard
	Score     float64
}

type output struct {
	world World
	steps int
}

func (p Program) Run(w World) output {
	steps := 0
	var label Label = 0
	var i Instruction
loop:
	for {
		steps++
		i = p[label]
		/* fmt.Println("step: ", steps, "i:", i, "inst:", i.inst, "world:", w) */
		switch i.inst {
		case END:
			break loop
		case INC:
			w[i.reg]++
			label = i.next
		case DEC:
			if w[i.reg] > 0 {
				w[i.reg]--
				label = i.next
			} else {
				label = i.jump
			}
		}
	}
	return output{w, steps}
}

var p Program = Program{0: Instruction{inst: DEC, reg: 0, next: 1, jump: 2}, 1: Instruction{inst: INC, reg: 1, next: 0}}

var w2 World = World{0: 10, 1: 5}
var w1 World = World{0: 10}

func main() {
	fmt.Println("test")

	for _, e := range r {
		o := p0.Run(e.in)
		fmt.Println("Input:", e.in)
		fmt.Println("output:", o)
		if o.world == e.out {
			fmt.Println("MATCH")
		} else {
			fmt.Println("NO MATCH")
		}
	}
	/* ochan := make(chan output, 1) */
	/* go func() { ochan <- p.Run(w1) }() */
	/* select { */
	/* case o := <-ochan: */
	/* 	fmt.Println("got output:", o) */
	/* case <-time.After(TimeOut * time.Second): */
	/* 	fmt.Println("BAD NEWS") */
	/* } */

}
