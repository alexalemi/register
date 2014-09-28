package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

type Instruction int

var Seperator rune = ':'

const (
	UNK Instruction = iota
	INC
	DEB
	END
)

type Label string
type Register string

type Statement struct {
	Label  Label
	Inst   Instruction
	Reg    Register
	Next   Label
	Branch Label
}

func (i Instruction) String() (s string) {
	switch i {
	case INC:
		s = "+"
	case DEB:
		s = "-"
	case END:
		s = "◼"
	case UNK:
		s = "?"
	}
	return
}

func (s Statement) String() (r string) {
	switch s.Inst {
	case INC:
		r = fmt.Sprintf("[%v:%v(%v)→%v]", s.Label, s.Inst, s.Reg, s.Next)
	case DEB:
		r = fmt.Sprintf("[%v:%v(%v)→%v↛%v]", s.Label, s.Inst, s.Reg, s.Next, s.Branch)
	case UNK, END:
		r = fmt.Sprintf("[%v:%v]", s.Label, s.Inst)
	}
	return
}

type Program map[Label]Statement
type World map[Register]int

func (p Program) String() string {
	s := ""
	for key, val := range p {
		if key != "" && key != "start" && val.Inst != UNK && key != "end" {
			s += fmt.Sprintf("%v", val)
		}
	}
	return s
}

func NewInstruction(s string) (inst Instruction) {
	switch s {
	case "deb", "dec", "-":
		inst = DEB
	case "inc", "+":
		inst = INC
	case "stop", "halt", "end", "◼":
		inst = END
	default:
		inst = UNK
	}
	return
}

func NewStatement(line string, lineno int) (stmt Statement, ok bool) {
	var s scanner.Scanner
	reader := strings.NewReader(line)
	s.Init(reader)
	s.Mode = scanner.ScanIdents | scanner.SkipComments
	ok = true

	// Try to get Label
	s.Scan()
	if s.TokenText() == "init" {
		ok = false
		return
	}
	// Check to see if there is a : next
	if s.Peek() == Seperator {
		stmt.Label = (Label)(s.TokenText())
		s.Scan() // consume :
		s.Scan()
	} else {
		stmt.Label = (Label)(strconv.Itoa(lineno))
	}

	// Try to get instruction
	stmt.Inst = NewInstruction(s.TokenText())

	// Try to get the Register
	s.Scan()
	stmt.Reg = (Register)(s.TokenText())

	// Try to get the Next
	s.Scan()
	stmt.Next = (Label)(s.TokenText())

	// If we have a deb, get the branch
	if stmt.Inst == DEB {
		s.Scan()
		stmt.Branch = (Label)(s.TokenText())
	}

	if s.ErrorCount > 0 {
		ok = false
	}
	return
}

func NewWorld(line string) (w World, err error) {
	w = make(map[Register]int)
	var s scanner.Scanner
	reader := strings.NewReader(line)
	s.Init(reader)
	s.Mode = scanner.ScanIdents | scanner.SkipComments

	// Try to get Label
	tok := s.Scan()
	if s.TokenText() != "init" {
		err = errors.New("Not a valid initialization line, must start with init:")
		return
	}
	s.Mode = scanner.ScanChars | scanner.SkipComments
	tok = s.Scan()
	s.Mode = scanner.ScanIdents | scanner.SkipComments
	tok = s.Scan()

	for tok != scanner.EOF {
		reg := (Register)(s.TokenText())
		s.Mode = scanner.ScanChars | scanner.SkipComments
		tok = s.Scan()
		s.Mode = scanner.ScanInts | scanner.SkipComments
		tok = s.Scan()
		val, err := strconv.Atoi(s.TokenText())
		if err != nil {
			return w, err
		}
		w[reg] = val
		s.Mode = scanner.ScanIdents | scanner.SkipComments
		tok = s.Scan()
	}

	return
}

var endStatement = Statement{Label: (Label)("end"), Inst: END}

func NewProgram(src io.Reader) (p Program, w World) {
	p = make(map[Label]Statement)
	w = make(map[Register]int)
	var err error
	scanner := bufio.NewScanner(src)
	lineno := 0
	for scanner.Scan() {
		lineno++
		low := strings.ToLower(scanner.Text())
		statement, ok := NewStatement(low, lineno)
		if !ok {
			w, err = NewWorld(low)
			if err != nil {
				panic(fmt.Sprintf("Problem reading line %v: %v", lineno, low))
			}
		}
		// fmt.Println(statement)
		p[statement.Label] = statement
		if lineno == 1 {
			p["start"] = statement
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	p[endStatement.Label] = endStatement
	return
}

func Simulate(p Program, winit *World) (w World, err error) {
	w = *winit
	label := (Label)("start")
	for {
		stmt := p[label]

		// fmt.Println("Currently on statement: ", stmt)
		switch stmt.Inst {
		case INC:
			if _, ok := w[stmt.Reg]; ok {
				w[stmt.Reg]++
			} else {
				w[stmt.Reg] = 1
			}
			// fmt.Printf("Increment %v to %v\n", stmt.Reg, w[stmt.Reg])
			label = stmt.Next
		case DEB:
			if val, ok := w[stmt.Reg]; ok && val > 0 {
				w[stmt.Reg]--
				label = stmt.Next
				// fmt.Printf("Decrement %v to %v\n", stmt.Reg, w[stmt.Reg])
			} else {
				label = stmt.Branch
				// fmt.Printf("Branch since %v is empty\n", stmt.Reg)
			}
		case END:
			return w, nil
			break
		case UNK:
			err := errors.New(fmt.Sprintf("Reached unknown instruction"))
			return w, err
		}
	}

	return
}

func main() {
	p, winit := NewProgram(os.Stdin)
	// fmt.Println("PROGRAM:", p)
	fmt.Println("INIT:", winit)

	w, err := Simulate(p, &winit)
	fmt.Println("OUT:", w)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
