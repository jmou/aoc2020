package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type parser struct {
	in []byte
}

func NewParser(in []byte) parser {
	// no built-in reverse
	out := make([]byte, len(in))
	for i, ch := range in {
		out[len(out)-i-1] = ch
	}
	return parser{out}
}

func (p *parser) readByte() (byte, error) {
	if len(p.in) == 0 {
		return 0, errors.New("unexpected END")
	}
	value := p.in[0]
	p.in = p.in[1:]
	return value, nil
}

func (p *parser) chompSpace() error {
	space, err := p.readByte()
	if err != nil {
		return err
	} else if space != ' ' {
		return errors.New("expected SPACE")
	}
	return nil
}

func (p *parser) parseNumber() (int, error) {
	// input numbers are all single digit
	num, err := p.readByte()
	return int(num - '0'), err
}

func (p *parser) ParseAtom() (int, error) {
	if len(p.in) == 0 {
		return 0, errors.New("unexpected END")
	} else if p.in[0] == ')' {
		p.in = p.in[1:]
		value, err := p.ParseExpr()
		if err != nil {
			return 0, err
		}
		paren, err := p.readByte()
		if err != nil {
			return 0, err
		}
		if paren != '(' {
			return 0, errors.New("expected (")
		}
		return value, nil
	} else {
		value, err := p.parseNumber()
		if err != nil {
			return 0, err
		}
		return value, nil
	}
}

func (p *parser) ParseExpr() (int, error) {
	right, err := p.ParseAtom()
	if err != nil {
		return 0, err
	}
	if len(p.in) == 0 || p.in[0] == '(' {
		return right, nil
	}
	if err := p.chompSpace(); err != nil {
		return 0, err
	}
	op, err := p.readByte()
	if err != nil {
		return 0, err
	}
	if err := p.chompSpace(); err != nil {
		return 0, err
	}
	left, err := p.ParseExpr()
	if err != nil {
		return 0, err
	}
	switch op {
	case '+':
		return left + right, nil
	case '*':
		return left * right, nil
	default:
		return 0, errors.New("expected OPERATOR")
	}
}

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parser := NewParser(scanner.Bytes())
		value, err := parser.ParseExpr()
		if err != nil {
			panic(err)
		}
		sum += value
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println(sum)
}
