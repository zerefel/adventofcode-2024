package main

import (
	"fmt"
	"math"
	"strings"
)

type AntennaSymbol string

type AntennaMap struct {
	Antennas  map[AntennaSymbol][]*Antenna
	Antinodes map[string]Antinode
	Rows      int
	Cols      int
	input     [][]string
}

func createAntennaMap() *AntennaMap {
	return &AntennaMap{
		Antennas:  make(map[AntennaSymbol][]*Antenna),
		Antinodes: make(map[string]Antinode),
		Rows:      0,
		Cols:      0,
		input:     make([][]string, 0),
	}
}

func (am *AntennaMap) PrintAntennaMap() {
	for i, row := range am.input {
		for j, col := range row {
			key := fmt.Sprintf("%d,%d", j, i)
			if _, exists := am.Antinodes[key]; exists && col == "." {
				fmt.Print("#")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func (am *AntennaMap) AddAntenna(a *Antenna) {
	am.Antennas[a.Symbol] = append(am.Antennas[a.Symbol], a)
}

func (am *AntennaMap) IsWithinBounds(a Antinode) bool {
	return a.X >= 0 && a.Y >= 0 && a.X < am.Cols && a.Y < am.Rows
}

func (am *AntennaMap) FindAntinodes(symbol AntennaSymbol) {
	antennas := am.Antennas[symbol]
	for i := 0; i < len(antennas); i++ {
		a := antennas[i]
		for j := i + 1; j < len(antennas); j++ {
			b := antennas[j]
			xDistance := int(math.Abs(float64(a.X) - float64(b.X)))
			yDistance := int(math.Abs(float64(a.Y) - float64(b.Y)))
			if xDistance == 0 && yDistance == 0 {
				continue
			}
			a1 := Antinode{Symbol: symbol}
			b1 := Antinode{Symbol: symbol}
			if a.X < b.X {
				a1.X = a.X - xDistance
				b1.X = b.X + xDistance
			} else if a.X > b.X {
				a1.X = a.X + xDistance
				b1.X = b.X - xDistance
			} else {
				a1.X = a.X
				b1.X = b.X
			}
			if a.Y < b.Y {
				a1.Y = a.Y - yDistance
				b1.Y = b.Y + yDistance
			} else if a.Y > b.Y {
				a1.Y = a.Y + yDistance
				b1.Y = b.Y - yDistance
			} else {
				a1.Y = a.Y
				b1.Y = b.Y
			}
			if am.IsWithinBounds(a1) {
				key := fmt.Sprintf("%d,%d", a1.X, a1.Y)
				if _, exists := am.Antinodes[key]; !exists {
					am.Antinodes[key] = a1
				}
			}
			if am.IsWithinBounds(b1) {
				key := fmt.Sprintf("%d,%d", b1.X, b1.Y)
				if _, exists := am.Antinodes[key]; !exists {
					am.Antinodes[key] = b1
				}
			}
		}
	}
}

type Antenna struct {
	Symbol AntennaSymbol
	X      int
	Y      int
}

type Antinode struct {
	Symbol AntennaSymbol
	X      int
	Y      int
}

func CalculateAntinodeLocations(am *AntennaMap) int {
	for symbol := range am.Antennas {
		am.FindAntinodes(symbol)
	}
	return len(am.Antinodes)
}

func parseAntennaMap(cityMap []string) *AntennaMap {
	am := createAntennaMap()
	am.Rows = len(cityMap)
	am.input = make([][]string, am.Rows)
	for i, row := range cityMap {
		cols := strings.Split(row, "")
		am.Cols = len(cols)
		am.input[i] = make([]string, am.Cols)
		for j, col := range cols {
			am.input[i][j] = col
			if col != "." {
				symbol := AntennaSymbol(col)
				a := &Antenna{Symbol: symbol, X: j, Y: i}
				am.AddAntenna(a)
			}
		}
	}
	return am
}

func main() {
	testCityMap := []string{
		".A...........5........................pL..........",
		".................................p......L.........",
		"......................................L...........",
		".......................................C..........",
		"........v...................7...............C.....",
		"..................................p........L......",
		".................vA......3........................",
		".......A.....3....................................",
		"........................s....X3...................",
		"..A......5.................9....3.................",
		".......8...........s.........7.............C..m...",
		"................8......t........7.......9.........",
		"....................o......Z.............y........",
		"...............s.......Y.v.o......y....0..........",
		"..................................................",
		"..5................8.......................m...J..",
		"5...............................0....aX...........",
		".V............v.s.........Z.o..7....a.............",
		"2..........f...........P..............9...J.M.....",
		"...............f..........P.....V......y....1.J...",
		"...g...................o.......0l...........N..B..",
		"..................Y...............................",
		"......G...............f.....Z..t..............1...",
		"............G......Z......h................B....C.",
		".........w....h.Y....j............a........J..y...",
		".............P....z..........................1....",
		"w.......P...z...R......r8.........................",
		"........w.........................................",
		".................h.G.........m............BM......",
		"......4.....fa.................G...i....X......W..",
		"V........4..............................tW.9...i..",
		"............2h..............0.......tX...M........",
		".....z.........................l..................",
		".......2..........................................",
		"..r........................Y................W...i.",
		".......w.........q..................i.............",
		".........H.2....4.................................",
		"..........Q.....j.......M.....lrN.................",
		"..x...H.Q.......O.....c...........................",
		"....H.......................S.....................",
		".....................O..S.......6..........b......",
		"...c.......F...Q.j.........l....T.....R...........",
		"...........Q.F.......c.I.....1.........R....T.....",
		"............F........I.O......r..T.............b..",
		"..n.........q.........F.I..............T..b.......",
		".......n...........z..O....x.......N........b.....",
		".....S............................................",
		"..........q.........cS..x4I......6................",
		"..j.....gn.q.......x...................N...6......",
		"...........g..n................R......B...........",
	}

	parsedMapTest := parseAntennaMap(testCityMap)
	locationsTest := CalculateAntinodeLocations(parsedMapTest)
	fmt.Printf("Total antinode locations: %d\n", locationsTest)
	parsedMapTest.PrintAntennaMap()
}
