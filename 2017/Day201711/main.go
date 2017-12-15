package Day201711

import (
	"sort"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 11, "Hex Ed"
}

func travel(input string) (int, int) {
	directions := strings.Split(input, ",")

	x, y, max := 0, 0, 0

	for _, val := range directions {
		switch val {
		case "n":
			y++
		case "ne":
			x++
		case "se":
			x++
			y--
		case "s":
			y--
		case "sw":
			x--
		case "nw":
			x--
			y++
		}

		vals := []int{x, y, -(x + y)}

		sort.Ints(vals)

		if vals[2] > max {
			max = vals[2]
		}
	}

	vals := []int{x, y, -(x + y)}

	sort.Ints(vals)

	return vals[2], max
}

func (td dayEntry) PartOne(inputData string) (string, error) {
	distance, _ := travel(inputData)
	return strconv.Itoa(distance), nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	_, maxDistance := travel(inputData)
	return strconv.Itoa(maxDistance), nil
}
