package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Command int
type Register int
type Label int

const MaxReg = 16
const MaxLen = 128
const TimeOut = 100

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

func (w World) Copy() (o World) {
	for i := range w {
		o[i] = w[i]
	}
	return o
}

func (p Program) Run(inw World) output {
	w := inw.Copy()
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
	case <-time.After(TimeOut * time.Millisecond):
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

type specimen struct {
	prog     Program
	score    float64
	age      int
	children int
}

const mutrate = 1e-5

func (p Program) Evolve() Program {
	var o Program
	for j, inst := range p {
		o[j] = inst
		if rand.Float64() < mutrate {
			o[j].inst = (Command)(rand.Int63n(3))
		}
		if rand.Float64() < mutrate {
			o[j].reg = (Register)(rand.Int63n(MaxReg))
		}
		if rand.Float64() < mutrate {
			o[j].next = (Label)(rand.Int63n(MaxLen))
		}
		if rand.Float64() < mutrate {
			o[j].jump = (Label)(rand.Int63n(MaxLen))
		}
	}
	return o
}

const CandSize = 1000

type cands [CandSize]specimen

var Candidates cands

type scorestruct struct {
	i int
	s float64
}

func (c *cands) Len() int           { return len(c) }
func (c *cands) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c *cands) Less(i, j int) bool { return c[i].score < c[j].score }

func (p Program) Copy() Program {
	var o Program
	for i := range p {
		o[i] = p[i]
	}
	return o
}

func (c *cands) Generation() {
	sort.Sort(c)
	for i := 0; i < 10; i++ {
		c[i].age++
		for j := 1; j < 100; j++ {
			z := j*10 + i
			c[z] = specimen{prog: c[i].prog.Copy()}
			c[i].children++
			c[z].prog.Evolve()
		}
	}

	signal := make(chan scorestruct, CandSize)
	for i := range c {
		go func(i int) {
			s := c[i].prog.Score(recipe)
			signal <- scorestruct{i, s}
		}(i)
	}
	for _ = range c {
		s := <-signal
		c[s.i].score = s.s
	}
	sort.Sort(c)
}

const generations = 100

func (c Command) String() string {
	switch c {
	case END:
		return "END"
	case INC:
		return "INC"
	case DEC:
		return "DEB"
	}
	return "NIL"
}

func (p Program) String() string {
	s := ""
	for i := range p {
		s += fmt.Sprintf("%v: %s %v; %v %v\n", i, p[i].inst, p[i].reg, p[i].next, p[i].jump)
	}
	return s
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := range Candidates {
		Candidates[i] = specimen{prog: p0.Evolve().Copy()}
	}
	Candidates[0].prog = p0

	for i := 0; i < 1000; i++ {
		a := Candidates[i].prog.Score(recipe)
		fmt.Println(i, a)
	}

	// for gen := 0; gen < generations; gen++ {
	// 	Candidates.Generation()
	// 	fmt.Print("Generation: ", gen, " | ")
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Print(Candidates[i].score, ", ")
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("DONE")

	f, err := os.Create("/tmp/dat2")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 3; i++ {
		f.WriteString(fmt.Sprint("PROGRAM ", i, "age:", Candidates[i].age, " children: ", Candidates[i].children, " score: ", Candidates[i].score, "\n"))
		f.WriteString(fmt.Sprint(Candidates[i].prog))
		f.WriteString("\n")
	}
}
