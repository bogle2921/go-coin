package main

import (
	"coin/ecc"
	"fmt"
	"math/big"
	"math/rand"
)

func multiplySet(min int, max int, order int64) {
	k := rand.Intn(max-min) + min
	fmt.Printf("k is %d\n", k)

	el := ecc.NewElement(big.NewInt(order), big.NewInt(int64(k)))
	for i := 0; i < int(order); i++ {
		fmt.Printf("%d * %d is %v\n", k, i, el.Scale(big.NewInt(int64(i))))
	}
}

// should return value == 1
// k ^ (p-1) % p == 1
// Fermat's theorem
func powerSet() {
	orders := []int{7, 11, 19, 31, 37}
	for _, p := range orders {
		for i := 1; i < p; i++ {
			el := ecc.NewElement(big.NewInt(int64(p)), big.NewInt(int64(i)))
			fmt.Printf("FE %v, power of p-1 = %v\n", el, el.Pow(big.NewInt(int64(p-1))))
		}
	}
}

func main() {
	// f44 := ecc.NewElement(big.NewInt(57), big.NewInt(44))
	// f33 := ecc.NewElement(big.NewInt(57), big.NewInt(33))
	// res := f44.Plus(f33)
	// fmt.Printf("44 + 33 = %v\n", res)               // 20
	// fmt.Printf("negate of 44 = %v\n", res.Negate()) // 37
	// fmt.Printf("44 - 33 = %v\n", f44.Minus(f33))    // 11
	// fmt.Printf("33 - 44 = %v\n", f33.Minus(f44))    // 46

	// // sanity check (46 + 44) % 57 == 33
	// fmt.Printf("(46 + 44) mod 57 == %d\n", (46+44)%57) // 33

	// f46 := ecc.NewElement(big.NewInt(57), big.NewInt(46))
	// // fmt.Printf("46 + 44 = %v\n", f46.Plus(f44)) // 33

	// fmt.Printf("46 * 46 = %v\n", f46.Multiply(f46))       // 7
	// fmt.Printf("46 ^ 58 = %v\n", f46.Pow(big.NewInt(58))) // 7

	// // multiplySet(1, 18, 19)
	// //powerSet()
	// f2 := ecc.NewElement(big.NewInt(int64(19)), big.NewInt(int64(2)))
	// f7 := ecc.NewElement(big.NewInt(int64(19)), big.NewInt(int64(7)))
	// fmt.Printf("2 / 7 = %v\n", f2.Divide(f7)) // 3

	// check (-1, -1) on y^2 = x^3 + 5x + 7
	// ecc.NewECPoint(big.NewInt(int64(-1)), big.NewInt(int64(-1)), big.NewInt(int64(5)), big.NewInt(int64(7))) // on curve - works
	// fmt.Printf("Point (-1, -1) on y^2 = x^3 + 5x + 7\n")
	// // check (-1, -2) on y^2 = x^3 + 5x + 7
	// ecc.NewECPoint(big.NewInt(int64(-1)), big.NewInt(int64(-2)), big.NewInt(int64(5)), big.NewInt(int64(7))) // not on curve - panics
	// fmt.Printf("Point (-1, -2) on y^2 = x^3 + 5x + 7\n")
}
