package Day201918

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type object struct {
	pos    common.Point
	isKey  bool
	letter string
}

func (o object) String() string {
	text := "Door"

	if o.isKey {
		text = "Key"
	}

	return fmt.Sprintf("{%v %v %v}", o.pos, text, o.letter)
}

type player struct {
	pos       common.Point
	inventory []object
}

func getData(input string) ([][]bool, []object, player) {
	lines := strings.Split(input, "\n")
	tunnels := make([][]bool, len(lines))

	for i := range tunnels {
		tunnels[i] = make([]bool, len(lines[i]))
	}

	var objects []object
	var p player
	for y := range lines {
		for x := range lines[y] {

			if lines[y][x] == 35 {
				tunnels[y][x] = true
			} else if lines[y][x] >= 97 && lines[y][x] <= 122 {
				obj := object{
					pos:    common.Point{X: x, Y: y},
					isKey:  true,
					letter: strings.ToUpper(string(lines[y][x]))}
				objects = append(objects, obj)
			} else if lines[y][x] >= 65 && lines[y][x] <= 90 {
				obj := object{
					pos:    common.Point{X: x, Y: y},
					isKey:  false,
					letter: string(lines[y][x])}
				objects = append(objects, obj)
			} else if lines[y][x] == 64 {
				p = player{
					pos:       common.Point{X: x, Y: y},
					inventory: []object{},
				}
			}
		}
	}

	return tunnels, objects, p
}

func printTunnels(tunnels [][]bool, objects []object, player player) {
	for y := range tunnels {
		for x := range tunnels[y] {
			char := " "
			if tunnels[y][x] {
				char = "#"
			}

			if player.pos.X == x && player.pos.Y == y {
				char = "@"
			}

			for i := range objects {
				if objects[i].pos.X == x && objects[i].pos.Y == y {
					if objects[i].isKey {
						char = strings.ToLower(objects[i].letter)
						break
					}
					char = objects[i].letter
					break
				}
			}
			fmt.Print(char)
		}

		fmt.Print("\n")
	}
}

type found struct {
	depth int
	thing object
}

func (f found) String() string {
	return fmt.Sprintf("{%v %v}", f.depth, f.thing)
}

func hasObject(pos common.Point, list []object) (bool, object) {
	for i := range list {
		if pos.X == list[i].pos.X && pos.Y == list[i].pos.Y {
			return true, list[i]
		}
	}
	return false, object{}
}

func hasKey(input string, inv []object) bool {
	for i := range inv {
		if inv[i].letter == input && inv[i].isKey {
			return true
		}
	}
	return false
}

func transfer(o object, items []object, p player) ([]object, player) {

	return items, player
}

func getPossibleMoves(tunnels [][]bool, objects []object, pos common.Point, inv []object, depth int, visits map[common.Point]bool) []found {
	visits[pos] = true

	fl := []found{}
	for _, dir := range common.Cardinal4 {
		testPos := pos.Add(dir)

		if _, ok := visits[testPos]; ok {
			continue
		}

		if ok, obj := hasObject(testPos, objects); ok {
			if obj.isKey || hasKey(obj.letter, inv) {
				fl = append(fl, found{depth, obj})
			}
			continue
		}

		if tunnels[testPos.Y][testPos.X] {
			continue
		}

		fl = append(fl, getPossibleMoves(tunnels, objects, testPos, inv, depth+1, visits)...)
	}

	return fl
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
