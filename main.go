// nolint
package main

import (
	"fmt"
	"math/big"
)

const prec = 200

func Mul(r1, r2 big.Float) big.Float {
	var z *big.Float = big.NewFloat(1)
	z.SetPrec(prec)
	return *z.Mul(&r1, &r2)
}

func Quo(r1, r2 big.Float) big.Float {
	var z *big.Float = big.NewFloat(1)
	z.SetPrec(prec)
	return *z.Quo(&r1, &r2)
}

func Add(r1, r2 big.Float) big.Float {
	var z *big.Float = big.NewFloat(1)
	z.SetPrec(prec)
	return *z.Add(&r1, &r2)
}

func Sub(r1, r2 big.Float) big.Float {
	var z *big.Float = big.NewFloat(1)
	z.SetPrec(prec)
	return *z.Sub(&r1, &r2)
}

func Sqrt(r1 big.Float) big.Float {
	var z *big.Float = big.NewFloat(1)
	z.SetPrec(prec)
	return *z.Sqrt(&r1)
}

func Pow2(r1 big.Float) big.Float {
	var z *big.Float = big.NewFloat(1)
	z.SetPrec(prec)
	return *z.Mul(&r1, &r1)
}

func newFloat(val float64) big.Float {
	result := *big.NewFloat(val)
	result.SetPrec(prec)
	return result
}

func main() {

	one, two, three, four, five, six, seven := newFloat(1), newFloat(2), newFloat(3), newFloat(4), newFloat(5), newFloat(6), newFloat(7)
	_, _, _, _, _, _, _ = one, two, three, four, five, six, seven
	r := newFloat(1) // unit circle radius
	A := newFloat(2) // A0 = area of the inner square

	// number of segments
	n := newFloat(3) // 3 = triangle, 4 = square, 5 = pentagon

	// length side
	Ln := Quo(three, Sqrt(three)) // triangle
	//Ln := Sqrt(two) // square
	//Ln := Mul(two, Sqrt(Sub(one, Pow2(Quo(Add(one, Sqrt(five)), four))))) // pentagon

	// inner length
	halfLn := Quo(Ln, two)
	Li := Sqrt(Sub(r, Pow2(halfLn)))
	//Li := Quo(one, two) // triangle
	//Li = Quo(one, Sqrt(two)) // square

	// area of the big curve to the area of the under triangle segment
	//curve := Quo(six, seven)
	//curve := Quo(one, two)
	curve := Quo(two, three)

	for i := 0; i < 100; i++ {

		// height
		h := Sub(r, Li)

		Atriangle := Mul(halfLn, h)
		Atriangles := Mul(Atriangle, n)
		Aarcs := Mul(Atriangles, curve)

		A = Add(A, Atriangles)
		AIterationRes := Add(A, Aarcs)

		if i >= 0 {
			i1 := Sqrt(Sub(one, Quo(Pow2(Ln), four)))
			halfMinusCurve := Sub(Quo(one, two), curve)
			i2 := Mul(halfMinusCurve, i1)
			i3 := Mul(curve, Ln)
			i4 := Add(i3, i2)
			res2 := Mul(n, i4)
			_ = res2

			i5 := Quo(Mul(Pow2(Ln), Ln), Mul(three, four))
			i6 := Quo(Ln, two)
			i7 := Quo(Pow2(Ln), Pow2(four))
			res3 := Mul(n, Sub(Add(i6, i5), i7))
			//res3 := Mul(n, i5)

			LnSqrdDivH := Quo(Pow2(Ln), h)

			fmt.Printf("i: %v, n %v, res %v, res3 %v, Ln^2/h %v\n", i, &n, &AIterationRes, &res3, &LnSqrdDivH)
		}

		// set the length and number of segments for next time
		n = Mul(n, two)

		Ln = Sqrt(Add(Pow2(h), Mul(halfLn, halfLn)))

		halfLn = Quo(Ln, two)
		Li = Sqrt(Sub(r, Pow2(halfLn)))
		//fmt.Printf("debug Li: %v\n", &Li)
	}
}
