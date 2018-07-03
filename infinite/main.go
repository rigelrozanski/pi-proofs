package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	r := newFloat(1) // unit circle radius
	n := newFloat(4) // number of segments
	Ln := Sqrt(two)  // length side
	halfLn := Quo(Ln, two)
	s := Sqrt(Sub(r, Pow2(halfLn))) // inner length

	for i := 0; i < 30; i++ {
		h := Sub(r, s) // height
		LnPow2DivH := Quo(Pow2(Ln), h)

		if i > 0 {
			Pi1 := Quo(Mul(Ln, n), two)

			LnFromPi := Quo(newFloat(math.Pi), n)
			Ln2, n2 := getNextLength(i, n, LnFromPi)
			for j := 1; j < 10; j++ {
				Ln2, n2 = getNextLength(i, n2, Ln2)
			}
			Pi2 := Quo(Mul(Ln2, n2), Pow(two, 9)) //????????? why four?

			fmt.Printf("i: %v, n %v, Ln^2/h %v, Pi1 %v, Pi2 %v\n", i, &n, &LnPow2DivH, &Pi1, &Pi2)
		}

		// set the length and number of segments for next time
		n = Mul(n, two)
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
