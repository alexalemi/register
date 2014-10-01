package main

import (
	"fmt"
	"math"
	"time"
)

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

// Run runs a world on the given initial condition
func (p Program) RunExp(w World) (o output, timeout bool) {
	ochan := make(chan output, 1)
	var out output
	go func() { ochan <- p.Run(w) }()
	select {
	case o := <-ochan:
		out = o
	case <-time.After(TimeOut * time.Second):
		timeout = true
	}
	return out, timeout
}

func ComputeScore(exp World, out output, timeout bool) float64 {
	s := 0.0
	if timeout {
		s += 1e5
	}
	s += float64(out.steps)
	for i, v := range exp {
		s += 100.0 * float64(math.Abs(float64(v)-float64(out.world[i])))
	}
	return s
}

type scoretup struct {
	exp     World
	out     output
	timeout bool
}

// Score scores a program
func (p Program) Score(r Recipe) float64 {
	N := len(r)
	reschan := make(chan scoretup, N+1)

	for _, exp := range r {
		go func() {
			o, t := p.RunExp(exp.in)
			reschan <- scoretup{exp.out, o, t}
		}()
	}

	score := 0.0
	for _ = range r {
        res := <-reschan
		score += ComputeScore(res.exp, res.out, res.timeout)
	}

	return score
}

func main() {
	fmt.Println("test")

    a := p0.Score(recipe)
    fmt.Println(a)

}
