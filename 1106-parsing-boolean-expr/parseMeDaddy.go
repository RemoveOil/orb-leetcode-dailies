package parsingbooleanexpr

import (
	"bytes"
	"fmt"
	"slices"
)

func parseBoolExpr(expression string) bool {
	parser := booleanParser{chars: bytes.NewBufferString(expression).Bytes()}

	res, err := parser.parseExpr()
	if err != nil {
		panic(fmt.Sprintf("You fucked up chicken. %v", err))
	}
	return res
}

type booleanParser struct {
	chars       []byte
	currentIndx int
	operators   []byte
	values      [][]bool
}

func (p *booleanParser) hasNext() bool   { return p.currentIndx < len(p.chars) }
func (p *booleanParser) currToken() byte { return p.chars[p.currentIndx] }

// Creates a new stack level
func (p *booleanParser) pushOp(op byte) {
	p.values = append(p.values, []bool{})
	p.operators = append(p.operators, op)
}
func (p *booleanParser) pushVal(val bool) {
	p.values[len(p.values)-1] = append(p.values[len(p.values)-1], val)
}

func (p *booleanParser) evaluateStack() (bool, error) {
	if len(p.operators) == 0 {
		return false, fmt.Errorf("evaluateStack: Operators empty. i: %v operands: %v", p.currentIndx, p.values)
	}
	var op byte
	var values []bool
	op, p.operators = pop(p.operators)
	values, p.values = pop(p.values)
	if len(values) == 0 {
		return false, fmt.Errorf("evaluateStack: No opoerands. i: %v operator: %v", p.currentIndx, op)
	}

	switch op {
	case '&':
		return !slices.Contains(values, false), nil
	case '|':
		return slices.Contains(values, true), nil
	case '!':
		if len(values) != 1 {
			return false, fmt.Errorf("bad operands for NOT operator: %v", values)
		}
		return !values[0], nil
	case '_':
		if len(values) != 1 {
			return false, fmt.Errorf("bad operands for NO-OP operator: %v", values)
		}
		return values[0], nil
	default:
		return false, fmt.Errorf("cannot evaluate?! op: %v, i: %v", op, p.currentIndx)
	}
}

func (p *booleanParser) parseExpr() (bool, error) {
	p.pushOp('_') // push a no-op operator to handle "t", "f" gracefully
	for p.hasNext() {
		curr := p.currToken()
		switch curr {
		case ',', ' ':
			p.currentIndx++ // ignore all commas / empty spaces
		case 'f', 't':
			p.pushVal(curr == 't')
			p.currentIndx++
		case '!', '&', '|':
			p.pushOp(curr)
			nextPos := p.currentIndx + 1
			if p.chars[nextPos] != '(' {
				return false, fmt.Errorf("expected %v at pos %v. Found %v", rune('('), nextPos, rune(p.chars[nextPos]))
			}
			p.currentIndx += 2 // skip the ( char
		case '(':
			// this means we've run into a naked paranthesis in the wild
			// we'll add a no-op operator to handle gracefully
			p.currentIndx++
			p.pushOp('_')
			fmt.Println("This happened?!@#")
		case ')':
			res, err := p.evaluateStack()
			if err != nil {
				return false, err
			}
			p.currentIndx++ // ignore all commas / empty spaces
			p.pushVal(res)
		default:
			return false, fmt.Errorf("we shouldn't have got here. i: %v", p.currentIndx)
		}
	}

	if !(len(p.operators) == 1 && len(p.values) == 1 && len(p.values[0]) == 1) {
		return false, fmt.Errorf("bad end state. parser: %v", p)
	}
	return p.values[0][0], nil
}

func pop[T any](list []T) (T, []T) { return list[len(list)-1], list[0 : len(list)-1] }
