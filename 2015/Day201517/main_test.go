package Day201517

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	matchChan := make(chan []int)

	go combinator(getData(Entry.PuzzleInput()), make([]int, 0), 0, matchChan)

	count := 0
	for combo := range matchChan {
		count++

		if count%1000 == 0 {
			fmt.Println(combo)
		}

	}
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	jars := getData(Entry.PuzzleInput())
	fmt.Println(jars[:3])
	fmt.Println(jars[3+1:])
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
