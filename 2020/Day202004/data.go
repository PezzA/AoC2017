package Day202004

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2020, 04, "Getting the boilerplate in place"
}

func (td dayEntry) PuzzleInput() string {
	return ``
}
