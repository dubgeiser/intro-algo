package main

/*
For each function f(n) and time t, determine the largest size n of a problem that
can be solved in time t, assuming that the algorithm to solve the problem takes
f(n) microseconds

Solution
========
Use the inverted version of the given function, this way max n = f-1(runtime)
For instance:
	f(n) -> sqrt(n) => max(n) = T^2
	f(n) -> n => max(n) == T
	...

fn(n)          fn(n)^-1
-----------------------
lg n           2^n
sqrt(n)        n^2
n              n
n lg n         (2^n)/n
n^2            sqrt(n)
n^3            sqrt(n, 3)?
2^n            lg n
n!             ???
*/

import (
	"fmt"
	"math"
)

func fac(n float64) float64 {
	f := float64(1)
	for i := n; i > 1; i-- {
		f = f * i
	}
	return f
}

func printResult(funcName, timeName string, n float64) {
	fmt.Printf("Max n for %s under 1 %s = %v\n", funcName, timeName, math.Trunc(n))
}

func main() {
	timeNames := []string{"second", "minute", "hour", "day", "month", "year", "century"}
	times := []float64{
		1_000_000,
		1_000_000 * 60,
		1_000_000 * 60 * 60,
		1_000_000 * 60 * 60 * 24,
		1_000_000 * 60 * 60 * 24 * 30,
		1_000_000 * 60 * 60 * 24 * 365,
		1_000_000 * 60 * 60 * 24 * 365 * 100,
	}

	// The given functions with corresponding inverse implementations
	funcNames := []string{"lg n", "sqrt(n)", "n", "n lg n", "n^2", "n^3", "2^n", "n!"}
	functions := []func(float64) float64{
		func(n float64) float64 { return math.Exp2(n) },
		func(n float64) float64 { return math.Pow(n, 2) },
		func(n float64) float64 { return n },
		func(n float64) float64 { return math.Exp2(n) / n },
		func(n float64) float64 { return math.Sqrt(n) },
		func(n float64) float64 { return math.Cbrt(n) },
		func(n float64) float64 { return math.Log2(n) },
	}

	for iF, f := range functions {
		for iT, t := range times {
			printResult(funcNames[iF], timeNames[iT], f(t))
		}
	}

	// Brute forcing the factorial function.
	// We could code up an inverted factorial, but it's gonna be a bit cumbersome
	// I think; the timings aren't necessarily factorials, so we need to find
	// the biggest lower one, etc...
	// I think we can be pretty sure that there won't be much elements that can
	// be handled in the given timings, so just check it... I think 100 elements
	// will not even be handled in a hundred years (Spoiler: they don't ;-) )
	for iT, t := range times {
		for n := 1; n < 100; n++ {
			if fac(float64(n)) > t {
				// If fac(n) has become larger than our timing, that means the
				// previous n was the largest n and we can go on to the next
				// timing.
				printResult("n!", timeNames[iT], float64(n-1))
				break
			}
		}
	}
}
