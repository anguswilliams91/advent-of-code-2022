package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"
)

type Solution[T, U any] struct {
	PartOne T
	PartTwo U
}

type intSlice []int64

// LoadInput loads a puzzle's input from a path p into a string.
func LoadInput(p string) string {
	input, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal("Couldn't load puzzle input.")
	}
	return string(input)
}

// Timer returns a function that runs a solution up to 100 times
// and measures the mean runtime and the standard deviation.
func Timer[T, U any](name string, solution func(string) Solution[T, U]) func(string) {
	return func(input string) {
		var ns intSlice
		var totalDuration time.Duration
		for i := 0; i < 100; i++ {
			start := time.Now()
			s := solution(input)
			t := time.Since(start)
			if i == 0 {
				fmt.Println("Part one: ", s.PartOne)
				fmt.Println("Part two: ", s.PartTwo)
			}
			ns = append(ns, t.Nanoseconds())
			totalDuration += t
			if totalDuration > 10*time.Second {
				break
			}
		}
		fmt.Printf(
			"%s runtime from %v trials = %s +/- %s\n",
			name,
			len(ns),
			time.Duration(ns.avg())*time.Nanosecond,
			time.Duration(ns.std())*time.Nanosecond,
		)
	}
}

func (xs intSlice) avg() int64 {
	var t int64 = 0
	for _, x := range xs {
		t += x
	}
	return t / int64(len(xs))
}

func (xs intSlice) std() int64 {
	mean := xs.avg()
	sumSq := 0
	for _, x := range xs {
		sumSq += int(math.Pow(float64(x-mean), 2))
	}
	return int64(math.Sqrt(float64(sumSq) / float64(len(xs)-1)))
}
