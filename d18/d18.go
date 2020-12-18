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
	// out := make([]byte, len(in))
	// for i, ch := range in {
	// 	out[len(out)-i-1] = ch
	// }
	return parser{in}
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

func (p *parser) ParseAtom() (int, int, error) {
	if len(p.in) == 0 {
		return 0, 0, errors.New("unexpected END")
	} else if p.in[0] == '(' {
		p.in = p.in[1:]
		p1, p2, err := p.ParseExpr()
		if err != nil {
			return 0, 0, err
		}
		paren, err := p.readByte()
		if err != nil {
			return 0, 0, err
		}
		if paren != ')' {
			return 0, 0, errors.New("expected (")
		}
		return p1, p2, nil
	} else {
		value, err := p.parseNumber()
		if err != nil {
			return 0, 0, err
		}
		return value, value, nil
	}
}

func (p *parser) ParseExpr() (int, int, error) {
	var atomsP1, atomsP2 []int
	p1, p2, err := p.ParseAtom()
	if err != nil {
		return 0, 0, err
	}
	atomsP1 = append(atomsP1, p1)
	atomsP2 = append(atomsP2, p2)
	for len(p.in) > 0 && p.in[0] != ')' {
		if err := p.chompSpace(); err != nil {
			return 0, 0, err
		}
		op, err := p.readByte()
		if err != nil {
			return 0, 0, err
		}
		if err := p.chompSpace(); err != nil {
			return 0, 0, err
		}
		atomsP1 = append(atomsP1, int(op))
		atomsP2 = append(atomsP2, int(op))
		p1, p2, err := p.ParseAtom()
		if err != nil {
			return 0, 0, err
		}
		atomsP1 = append(atomsP1, p1)
		atomsP2 = append(atomsP2, p2)
	}

	p1 = atomsP1[0]
	for i := 1; i < len(atomsP1); i += 2 {
		switch atomsP1[i] {
		case '+':
			p1 += atomsP1[i+1]
		case '*':
			p1 *= atomsP1[i+1]
		default:
			if len(atomsP1) > 1 {
				return 0, 0, errors.New("unknown operator")
			}
		}
	}

	// answer too low
	j := 1
	for i := 1; i < len(atomsP2); i += 2 {
		if atomsP2[i] == '+' {
			atomsP2[j-1] = atomsP2[j-1] + atomsP2[i+1]
		} else {
			atomsP2[j] = atomsP2[i]
			atomsP2[j+1] = atomsP2[i+1]
			j += 2
		}
	}
	atomsP2 = atomsP2[:j]

	p2 = 1
	// operators must all be '*'
	for i := 0; i < len(atomsP2); i += 2 {
		p2 *= atomsP2[i]
	}

	return p1, p2, nil
}

func main() {
	sumP1, sumP2 := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parser := NewParser(scanner.Bytes())
		p1, p2, err := parser.ParseExpr()
		if err != nil {
			panic(err)
		}
		sumP1 += p1
		sumP2 += p2
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println(sumP1)
	fmt.Println(sumP2)
}
