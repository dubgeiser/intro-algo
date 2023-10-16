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

func printTableHeader(timeNames []string) {
	fmt.Printf("%-7s", " ")
	for _, t := range timeNames {
		fmt.Printf("%16s", t)
	}
	fmt.Println()
}

func printFuncBegin(funcName string) {
	fmt.Printf("%-7s", funcName)
}

func printMaxN(n float64) {
	fmt.Printf("%16v", math.Trunc(n))
}

func printFuncEnd() {
	fmt.Println()
}

func main() {
	timeNames := []string{"1 second", "1 minute", "1 hour", "1 day", "1 month", "1 year", "1 century"}
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

		// float64 isn't enough, yields +Inf :-(
		// But resolvable if you write it as 2^T
		func(n float64) float64 { return math.Exp2(n) },

		func(n float64) float64 { return math.Pow(n, 2) },
		func(n float64) float64 { return n },

		// Since math.Exp2(n) yields +Inf, dividing by whatever yields +Inf :-(
		// But resolvable if you write it as (2^T)/T
		func(n float64) float64 { return math.Exp2(n) / n },

		func(n float64) float64 { return math.Sqrt(n) },
		func(n float64) float64 { return math.Cbrt(n) },
		func(n float64) float64 { return math.Log2(n) },
	}

	printTableHeader(timeNames)
	for iF, f := range functions {
		printFuncBegin(funcNames[iF])
		for _, t := range times {
			printMaxN(f(t))
		}
		printFuncEnd()
	}

	// Brute forcing the factorial function.
	// We could code up an inverted factorial, but it's gonna be a bit cumbersome
	// I think; the timings aren't necessarily factorials, so we need to find
	// the biggest lower one, etc...
	// I think we can be pretty sure that there won't be much elements that can
	// be handled in the given timings, so just check it... I think 100 elements
	// will not even be handled in a hundred years (Spoiler: they don't ;-) )
	printFuncBegin("n!")
	for _, t := range times {
		for n := 1; n < 100; n++ {
			if fac(float64(n)) > t {
				// If fac(n) has become larger than our timing, that means the
				// previous n was the largest n and we can go on to the next
				// timing.
				printMaxN(float64(n - 1))
				break
			}
		}
	}
	printFuncEnd()
}
