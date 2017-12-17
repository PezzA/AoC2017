package Day201721

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 21, "Getting the boilerplate in place"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PuzzleInput() string {
	return "Actual Data"
}
