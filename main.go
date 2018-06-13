package main

import (
	"fmt"
	"math/big"
)

func main() {

	r := big.NewFloat(1) // unit circle radius

	A := big.NewFloat(2) // A0 = area of the inner square

	// length of the inner side (start with the side length of the square)
	var z, z2, z3 *big.Float = big.NewFloat(1), big.NewFloat(1), big.NewFloat(1)
	two := big.NewFloat(2)
	sqrt2 := z.Sqrt(two)
	Ls := sqrt2 // length side

	// inner length
	Li := z.Quo(big.NewFloat(1), sqrt2) // 1/sqrt(2)

	// area of the big curve to the area of the under triangle segment
	curve := z.Quo(big.NewFloat(5), big.NewFloat(7))

	// number of segments
	n := big.NewFloat(4)

	for i := 0; i < 2; i++ {
		fmt.Printf("debug i: %v\n", i)

		// height
		h := z.Sub(r, Li)

		halfLs := z.Quo(Ls, two)
		fmt.Printf("debug halfLs: %v\n", halfLs)
		Atriangle := z.Quo(halfLs, h)
		fmt.Printf("debug Atriangle: %v\n", Atriangle)
		Atriangles := z.Mul(Atriangle, n)
		fmt.Printf("debug Atriangles: %v\n", Atriangles)
		Aarcs := z.Mul(Atriangles, curve)
		fmt.Printf("debug Aarcs: %v\n", Aarcs)

		A = z.Add(A, Atriangles)
		AIterationRes := z.Add(A, Aarcs)

		fmt.Printf("i: %v, res %v\n", i, AIterationRes)

		// set the length and number of segments for next time
		n = n.Mul(n, two)

		z = z.Mul(h, h)
		z2 = z2.Mul(halfLs, halfLs)
		z3 = z3.Add(z, z2)
		Ls = z.Sqrt(z3)

		halfLs = z.Quo(Ls, two)
		z2 = z2.Mul(halfLs, halfLs)
		fmt.Printf("debug z2: %v\n", z2)
		z3 = z3.Sub(r, z2)
		fmt.Printf("debug z3: %v\n", z3)
		Li = z.Sqrt(z3)
	}
}
