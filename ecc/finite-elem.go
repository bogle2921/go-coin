package ecc

import (
	"fmt"
	"math"
)

type FieldElement struct {
	order uint64
	value uint64
}

func NewElement(order uint64, value uint64) *FieldElement {
	if value >= order {
		err := fmt.Sprintf("Value must be between 0 and %d\n", order-1)
		panic(err)
	}

	return &FieldElement{
		order: order,
		value: value,
	}
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement {order: %d, value: %d}", f.order, f.value)
}

func (f *FieldElement) Equals(other *FieldElement) bool {
	return f.order == other.order && f.value == other.value
}

func (f *FieldElement) assertOrder(other *FieldElement) {
	if f.order != other.order {
		panic("Must add elements with same order")
	}
}
func (f *FieldElement) Plus(other *FieldElement) *FieldElement {
	f.assertOrder(other)

	// need to mod on order
	return NewElement(f.order, (f.value+other.value)%f.order)
}

// given a + b == 0, b == -a
// (a + b) % order == 0 -> b = order - a
func (f *FieldElement) Negate() *FieldElement {
	return NewElement(f.order, f.order-f.value)
}

// c = a - b
// (b + c) % order = a -> (a + (-b)) % order
func (f *FieldElement) Minus(other *FieldElement) *FieldElement {
	return f.Plus(other.Negate())
}

func (f *FieldElement) Multiply(other *FieldElement) *FieldElement {
	f.assertOrder(other)
	return NewElement(f.order, (f.value*other.value)%f.order)
}

func (f *FieldElement) Pow(power int64) *FieldElement {
	return NewElement(f.order, uint64(math.Pow(float64(f.value), float64(power)))%f.order)
}

func (f *FieldElement) Scale(val uint64) *FieldElement {
	return NewElement(f.order, (f.value*val)%f.order)
}
