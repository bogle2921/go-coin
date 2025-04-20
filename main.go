package main

import (
	"coin/ecc"
	"fmt"
	"math/rand"
)

func multiplySet(min int, max int, order uint64) {
	k := rand.Intn(max-min) + min
	fmt.Printf("k is %d\n", k)

	el := ecc.NewElement(order, uint64(k))
	for i := 0; i < int(order); i++ {
		fmt.Printf("%d * %d is %v\n", k, i, el.Scale(uint64(i)))
	}
}
func main() {
	// f44 := ecc.NewElement(57, 44)
	// f33 := ecc.NewElement(57, 33)
	// res := f44.Plus(f33)
	// fmt.Printf("44 + 33 = %v\n", res)               // 20
	// fmt.Printf("negate of 44 = %v\n", res.Negate()) // 37
	// fmt.Printf("44 - 33 = %v\n", f44.Minus(f33))    // 11
	// fmt.Printf("33 - 44 = %v\n", f33.Minus(f44))    // 46

	// // sanity check (46 + 44) % 57 == 33
	// fmt.Printf("(46 + 44) mod 57 == %d\n", (46+44)%57) // 33

	// f46 := ecc.NewElement(57, 46)
	// fmt.Printf("46 + 44 = %v\n", f46.Plus(f44)) // 33

	// fmt.Printf("46 * 46 = %v\n", f46.Multiply(f46)) // 7
	// fmt.Printf("46 ^ 2 = %v\n", f46.Pow(2))         // 7

	multiplySet(1, 18, 19)
}
