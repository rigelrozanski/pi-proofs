// nolint
package main

import (
	"fmt"
	"math/big"
)

const prec = 100

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

	one, two, three, four, six, seven := newFloat(1), newFloat(2), newFloat(3), newFloat(4), newFloat(6), newFloat(7)
	_, _, _, _, _, _ = one, two, three, four, six, seven
	r := newFloat(1) // unit circle radius
	A := newFloat(2) // A0 = area of the inner square

	// length of the inner side (start with the side length of the square)
	sqrt2 := Sqrt(two)
	Ls := sqrt2 // length side
	//fmt.Printf("debug Ls: %v\n", &Ls)

	// inner length
	Li := Quo(one, sqrt2) // 1/sqrt(2)
	//fmt.Printf("debug Li: %v\n", &Li)

	// area of the big curve to the area of the under triangle segment
	//curve := Quo(six, seven)
	curve := Quo(two, three)
	//curve := Quo(one, two)

	// number of segments
	n := newFloat(4)

	for i := 0; i < 1000; i++ {
		//fmt.Printf("debug i: %v\n", i)

		// height
		h := Sub(r, Li)
		//fmt.Printf("debug h: %v\n", &h)

		halfLs := Quo(Ls, two)
		//fmt.Printf("debug halfLs: %v\n", &halfLs)
		Atriangle := Mul(halfLs, h)
		//fmt.Printf("debug Atriangle: %v\n", &Atriangle)
		Atriangles := Mul(Atriangle, n)
		//fmt.Printf("debug Atriangles: %v\n", &Atriangles)
		Aarcs := Mul(Atriangles, curve)
		//fmt.Printf("debug Aarcs: %v\n", &Aarcs)

		A = Add(A, Atriangles)
		AIterationRes := Add(A, Aarcs)

		if i > 10 {
			i1 := Sqrt(Sub(one, Quo(Pow2(Ls), four)))
			halfMinusCurve := Sub(Quo(one, two), curve)
			i2 := Mul(halfMinusCurve, i1)
			i3 := Mul(curve, Ls)
			i4 := Add(i3, i2)
			res2 := Mul(n, i4)

			i5 := Quo(Ls, two)
			i6 := Mul(i5, i1)
			i7 := Quo(Pow2(four), Mul(three, Ls))
			res3 := Mul(n, Add(i6, i7))

			fmt.Printf("i: %v, n %v, res %v, res2 %v, res3 %v\n", i, &n, &AIterationRes, &res2, &res3)
		}

		// set the length and number of segments for next time
		n = Mul(n, two)

		Ls = Sqrt(Add(Pow2(h), Mul(halfLs, halfLs)))

		halfLs = Quo(Ls, two)
		Li = Sqrt(Sub(r, Pow2(halfLs)))
		//fmt.Printf("debug Li: %v\n", &Li)
	}
}
