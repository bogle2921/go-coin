package ecc

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	order *big.Int
	value *big.Int
}

func NewElement(order *big.Int, value *big.Int) *FieldElement {
	if order.Cmp(value) == -1 {
		var op big.Int
		err := fmt.Sprintf("Value must be between 0 and %d\n", op.Sub(order, big.NewInt(1)))
		panic(err)
	}

	return &FieldElement{
		order: order,
		value: value,
	}
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement {order: %s, value: %s}", f.order.String(), f.value.String())
}

func (f *FieldElement) Equals(other *FieldElement) bool {
	return f.order.Cmp(other.order) == 0 && f.value.Cmp(other.value) == 0
}

func (f *FieldElement) assertOrder(other *FieldElement) {
	if f.order.Cmp(other.order) != 0 {
		panic("Must add elements with same order")
	}
}
func (f *FieldElement) Plus(other *FieldElement) *FieldElement {
	f.assertOrder(other)

	// need to mod on order
	var op big.Int
	return NewElement(f.order, op.Mod(op.Add(f.value, other.value), f.order))
}

// given a + b == 0, b == -a
// (a + b) % order == 0 -> b = order - a
func (f *FieldElement) Negate() *FieldElement {
	var op big.Int
	return NewElement(f.order, op.Sub(f.order, f.value))
}

// c = a - b
// (b + c) % order = a -> (a + (-b)) % order
func (f *FieldElement) Minus(other *FieldElement) *FieldElement {
	return f.Plus(other.Negate())
}

func (f *FieldElement) Multiply(other *FieldElement) *FieldElement {
	f.assertOrder(other)
	var op big.Int
	mul := op.Mul(f.value, other.value)
	return NewElement(f.order, op.Mod(mul, f.order))
}

func (f *FieldElement) Pow(power *big.Int) *FieldElement {
	var op big.Int
	// This can be optimized since k^(p-1) % p = 1, power > p-1 => power%(p-1)
	t := op.Mod(power, op.Sub(f.order, big.NewInt(int64(1))))
	res := op.Exp(f.value, t, nil)
	mod := op.Mod(res, f.order)
	return NewElement(f.order, mod)
}

func (f *FieldElement) Scale(val *big.Int) *FieldElement {
	var op big.Int
	res := op.Mul(f.value, val)
	res = op.Mod(res, f.order)
	return NewElement(f.order, res)
}

// since k ^ (p-1) % p = 1, k ^ (p-2) = inverse k => a / b = a . b^(p-2)
func (f *FieldElement) Divide(other *FieldElement) *FieldElement {
	f.assertOrder(other)
	var op big.Int
	otherRev := other.Pow(op.Sub(f.order, big.NewInt(int64(2))))
	return f.Multiply(otherRev)
}
