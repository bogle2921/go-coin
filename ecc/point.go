package ecc

import (
	"fmt"
	"math/big"
)

type OPERATION int

const (
	ADD OPERATION = iota
	SUB
	MUL
	DIV
	POW
)

type Point struct {
	// coefficients of y^2 = x^3 + ax + b
	a *big.Int
	b *big.Int

	x *big.Int
	y *big.Int
}

func DoOperation(x *big.Int, y *big.Int, oper OPERATION) *big.Int {
	var op big.Int
	switch oper {
	case ADD:
		return op.Add(x, y)
	case SUB:
		return op.Sub(x, y)
	case MUL:
		return op.Mul(x, y)
	case DIV:
		return op.Div(x, y)
	case POW:
		return op.Exp(x, y, nil)
	}
	panic("unknown operation")
}

func NewECPoint(x *big.Int, y *big.Int, a *big.Int, b *big.Int) *Point {
	// ensure point is on curve
	left := DoOperation(y, big.NewInt(int64(2)), POW)
	x3 := DoOperation(x, big.NewInt(int64(3)), POW)
	ax := DoOperation(a, x, MUL)
	right := DoOperation(DoOperation(x3, ax, ADD), b, ADD)
	if left.Cmp(right) != 0 {
		err := fmt.Sprintf("Point (%v, %v) not on curve with a: %v, b: %v\n", x, y, a, b)
		panic(err)
	}
	return &Point{
		a: a,
		b: b,
		x: x,
		y: y,
	}
}

func (p *Point) Equals(other *Point) bool {
	return p.a.Cmp(other.a) == 0 && p.b.Cmp(other.b) == 0 && p.x.Cmp(other.x) == 0 && p.y.Cmp(other.y) == 0
}

func (p *Point) NotEqual(other *Point) bool {
	return ! p.Equals(other)
}
