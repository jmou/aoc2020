package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int
}

type machine struct {
	instructions []instruction
	accumulator  int
	pc           int
	executed     map[int]bool
}

func newMachine(instructions []instruction) machine {
	return machine{instructions, 0, 0, make(map[int]bool)}
}

func (m *machine) Step() {
	m.executed[m.pc] = true
	in := m.instructions[m.pc]
	switch in.op {
	case "acc":
		m.accumulator += in.arg
		m.pc++
	case "jmp":
		m.pc += in.arg
	case "nop":
		m.pc++
	}
}

func (m machine) LoopDetected() bool {
	return m.executed[m.pc]
}

func (m machine) Terminated() bool {
	return m.pc == len(m.instructions)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var instructions []instruction
	for {
		op, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		op = strings.Trim(op, " ")
		argString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		arg, err := strconv.Atoi(argString[1 : len(argString)-1])
		if err != nil {
			panic(err)
		}
		if argString[0] == '-' {
			arg *= -1
		}
		instructions = append(instructions, instruction{op, arg})
	}

	m := newMachine(instructions)
	for !m.LoopDetected() {
		m.Step()
	}
	fmt.Println(m.accumulator)

	for i, in := range instructions {
		switch in.op {
		case "jmp":
			instructions[i].op = "nop"
		case "nop":
			instructions[i].op = "jmp"
		default:
			continue
		}
		m = newMachine(instructions)
		for !m.LoopDetected() && !m.Terminated() {
			m.Step()
		}
		if m.Terminated() {
			fmt.Println(m.accumulator)
		}
		instructions[i].op = in.op
	}
}
