package Day202017

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`.#.
..#
###`)

	data.print(3, 3, 0, 0)

	Expect(len(data)).Should(Equal(3 * 3))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	for i := 1; i <= 6; i++ {
		data = data.cycle(8, 8, 0, i)

		fmt.Println(i, data.getActiveCubes())
	}

	Expect(data.getActiveCubes()).Should(Equal(112))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

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