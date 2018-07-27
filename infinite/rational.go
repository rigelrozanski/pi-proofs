//nolint
package main

import (
	"math/big"
)

var (
	one   = newFloat(1)
	two   = newFloat(2)
	three = newFloat(3)
	four  = newFloat(4)
	five  = newFloat(5)
	six   = newFloat(6)
	seven = newFloat(7)
)

const prec = 50

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

func Pow(r1 big.Float, power int) big.Float {
	for i := 1; i < power; i++ {
		r1 = Mul(r1, r1)
	}
	return r1
}

func newFloat(val float64) big.Float {
	result := *big.NewFloat(val)
	result.SetPrec(prec)
	return result
}
