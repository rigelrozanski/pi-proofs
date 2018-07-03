package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	r := newFloat(1)        // unit circle radius
	n := newFloat(4)        // number of segments
	nInt := 4               // number of segments
	Pi := newFloat(math.Pi) // number of segments
	Ln := Sqrt(two)         // length side
	halfLn := Quo(Ln, two)
	s := Sqrt(Sub(r, Pow2(halfLn))) // inner length

	for i := 0; i < 50; i++ {
		h := Sub(r, s) // height
		LnPow2DivH := Quo(Pow2(Ln), h)
		_ = LnPow2DivH

		// try to fast-forward a bunch, check this ratio

		//hj := Mul(h, one) // cheap copy method
		//nj := Mul(n, one)
		//Lnj := Mul(Ln, one)
		//halfLnj := Mul(halfLn, one)
		//sj := Mul(s, one)
		//for j := 0; j < 1000; j++ {
		//hj = Sub(r, sj) // height

		//// set the length and number of segments for next time
		//nj = Mul(nj, two)
		//Lnj = Sqrt(Add(Pow2(hj), Pow2(halfLnj)))

		//halfLnj = Quo(Lnj, two)
		//sj = Sqrt(Sub(r, Pow2(halfLnj)))
		//}
		//LnFaster := Mul(Lnj, Quo(nj, n))

		//  easy way just assume Pi
		LnFaster := Quo(Mul(two, Pi), n)
		ratioLn := Quo(LnFaster, Ln)
		ratioLn2 := Pow(ratioLn, i*2)

		if i > 0 {
			Pi1 := Quo(Mul(Ln, n), two)

			LnFromPi := Quo(newFloat(math.Pi), n)
			Ln2, n2 := getNextLength(i, n, LnFromPi)
			for j := 1; j < 10; j++ {
				Ln2, n2 = getNextLength(i, n2, Ln2)
			}
			Pi2 := Quo(Mul(Ln2, n2), Pow(two, 9)) //????????? why four?

			Pi2 = Quo(Mul(Add(Ln, Mul(two, h)), n), two)
			ratio := Quo(Pi1, Pi2)
			_ = ratio

			fmt.Printf("i: %v,\tn %v,\tPi1 %v,\tPi2 %v, \tLnF/Ln %v\n", i, &n, &Pi1, &Pi2, &ratioLn2)
		}

		// set the length and number of segments for next time
		n = Mul(n, two)
		nInt = nInt * 2
		Ln = Sqrt(Add(Pow2(h), Pow2(halfLn)))

		halfLn = Quo(Ln, two)
		s = Sqrt(Sub(r, Pow2(halfLn)))
	}
}

func getNextLength(maxIter int, n, prevLayerLn big.Float) (Ln, nOut big.Float) {
	Ln, n, _, _, _ = nextLength(newFloat(0), n, prevLayerLn, 0, maxIter)
	return Ln, n
}

// CONTRACT, computation and n must be passed in as Zeros
func nextLength(computation, n, prevLayerLn big.Float, iter, maxIter int) (computationO, nOut, prevLayerLnO big.Float, iterO, maxIterO int) {
	n = Mul(n, two)
	if iter == 0 { // if the first iteration
		computation = Sqrt(Sub(four, Pow2(prevLayerLn)))
		return nextLength(computation, n, prevLayerLn, iter+1, maxIter)
	}
	if iter == maxIter {
		Ln := Sqrt(Sub(two, computation))
		return Ln, n, prevLayerLn, iter, maxIter
	}
	computation = Sqrt(Add(two, computation))
	return nextLength(computation, n, prevLayerLn, iter+1, maxIter)
}
