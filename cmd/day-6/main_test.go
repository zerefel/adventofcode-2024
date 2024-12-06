package main

import "testing"

var sampleLab = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

var fullLab = []string{
	"...............................#.........#............#...#.............................#....................#.....#........##....",
	"..##..##.#.#....................##....#..#.....#...............#..#.....#.....................#...#.............#...#...#.........",
	"...................#............#..................................................#............................#..#..............",
	"............#...................................................#.........................#..#....#....................#..........",
	".#...........#...............................#..#....................................#............#.#............#................",
	".........................#..........................................#.................#...#.....#.#...............................",
	"..#..................#......................................#......#........#.......#..........#.........#..#.........#...........",
	"..........#..............#....................................#................................................#........#........#",
	"..#...............................................................................................................#.#..#..........",
	"....#....................##...................#..#....#....#.........#.......#.....................................#..............",
	"........#...................#.......#.........................................#.........#........#....................#...........",
	"................................................................#.......#......#...........#................#.....................",
	"...........................................#.............#...........##.#................................#...................#....",
	"............................#......#........#...#.................................................................................",
	".......#.............#...............................#.............................................#.............#.........#......",
	"......................................................#.........................#....#................................#...........",
	"...............#...........................................#..........#.........#..#..................#...........................",
	".....................................##...#................................#..#......................#.......#..........#.........",
	"..............#.......#....................#.....................................................#.......#...................##..#",
	"...............................................#.............#..#...#.....#..................................................#....",
	".....................................#........#...........................................................#......#................",
	"........##............#...............................#....#......................................................................",
	".....#............#...................................................................#.............................#.............",
	"........#...............#......#..........#........#.....#........#..........#...............................................#..#.",
	"..................#.............#........#......#.................................................#..........................#....",
	".................................#............#...............................................................#..##...............",
	"......................................#.#...........#........#.................................................................#..",
	"...........#...........#..............#..#.#...........#............................................#...#.....................#...",
	"....##......................#..........#..............#.............#.....#............................#....................#.....",
	"..................#...................#..........#............................................#...#.....................#.........",
	".....................#...........#.........................................................#......................................",
	".#..............#....#..#.#...##.........#....#.............#..........#.......#...#.............................#.........#...#..",
	"..................#...............#.....................#..#.....................#........#.......................................",
	"......#..#......#...#..................##.............#...................#.....#.....#................................#........#.",
	"....................#.#................#....#..................................................#............#................#....",
	"...........................#..#.#......#........................#......#..........................................................",
	"..#.#...............................................................#.#........................#.....................#..#.........",
	"...#..............................#......................................................#...#....................#.#.............",
	"....................#.......#...........#.....#.......................#........#...........................##.....................",
	"............#.........................................................#..............................................#.........#..",
	"..................#.#..............................#......#..................................#..........#......#..#...............",
	"....##................#.#..................#.....#...........................................#.......................##...........",
	"......#...........................#......#.......#..........#................#..#......#..........#..............#..#.............",
	"........#.................#...#.........................................#..........#.........#...................................#",
	".#..........#....#.....#.............#...........#..................#.................#...#.......................................",
	"...........................................................#......#...#...........................................##........#....#",
	"............#.............................................................................#..........................#............",
	".........#..............................................................................................#.........#.....#......#..",
	".........#.......................................................................#........................#..#.....#.#............",
	"............................#.............#..........#..#...........#......#..............#..........................#............",
	"..#.#............#.........#.........#..................#.......................#............................................#....",
	"...#......................#......................................................#...........#..........................#.........",
	".#....#...............................................................................................................#...........",
	".#..#................#................................................................#....^.................#.............#......",
	"........................................................................................#.....................#...................",
	"........##................#.............................#.........................................................................",
	".#.#.................#.....................#.....................#..........................#.................#...............#...",
	"........#....................................................................................................#....................",
	"##......#.............................#.......#........................................................................#.....##...",
	"..........#..........................................#.................#.....#............................................#.......",
	"......................#........................#.......................................#............#.............................",
	".............................#.#.......................................................................................#..........",
	"...#.........#.....#........................#.................#...........#.....................#.................................",
	".........#...........................................................................#...........................#................",
	"......#...#............#..........#........................#...................#...........#................................#.....",
	".....#...#....#..........................................................#......................#.....................#..........#",
	"#.............#....................#...............................#.......#.#.............#.....................#.............#..",
	"............................................................#.......#...................#.................................#......#",
	"...#..#......................................#.........................................#..#.#.......##.......#............#..#....",
	".....#.#...#...#.#................................................................................................................",
	"...............................................................................#........#...................#..#.........#....#...",
	"......#.........#.............#........................................................................................#.##.......",
	".......................#...................................................................................................#......",
	".........................................................#....................................#...................##.........#....",
	".#..............#.#.#.............................#........................#........................#...................#.#.......",
	".........#.....#..........................................................................#.................#.................#...",
	".........#..............#...............................#.........#......#.#......................................................",
	"......................................................................................................#....#............#.......#.",
	".....#.................#.......#.....#..................................#...............................................#.........",
	"...............................................................#.............#.................................................#..",
	"...#.....................................................................#...#.......#......#.....................................",
	".................................#...#.#............................................................#...#.........................",
	"...........#..#.................................#....#...................#.............#.........#................................",
	"..................................................#..#...........#...........................#....................................",
	".....#.....................................#......#......#...........................................................#............",
	"..............#..........................................#.........#.............#..#............................#................",
	"..........................................................................#..........#.......................#....................",
	".......##................#................................................................................................#.......",
	".#.....#.....#..................#............#......#............................#.........#......................................",
	"..##....#..................................................................................................................#......",
	".......#......................#..................#...#.....................................#............#..............#..........",
	"...............................#................#...#.................................##......................#...................",
	"...............#..............................................................................................#.....##.......#....",
	"................................#..........#........................................#..#..........................................",
	"..................................................#.................................................................#.............",
	"....................#....#.....................#.................................#.........#.....#......#......................#..",
	"................#.............#............#................#........................................................#..........#.",
	".#.......#......##...................#....#..............................................................#.................#......",
	".........#.#............................................................................#........................#................",
	"...............#........##.....#.....................................#..#...........#......#......................................",
	"......#.............................................###..#.........#..............#...........................#.......#...........",
	".....................................#.........#...........................................................##.....................",
	".#....#.......................#.........#................................................................................#..#.....",
	"....#..#.......................................................................................................##............#....",
	"..#...................................#....#......#....................#...............................#..........................",
	"...................#....##........................................##..........................................#...................",
	"..........#.........#.............................................................#..#.........#...#............#.......#..#......",
	"....#...........#...##.........................#......................#.........#..#..............................................",
	"....................#.....#.....................................#...............................................................#.",
	".........#.......................#.................#...#....................................................#.#......#........#...",
	".................#.................................#..#..............#..........#...#.........#........#.##....##.................",
	".............#...........................#......................#.......................#.........................................",
	"...............#.....#......#........#....#.......................................#.............................#.................",
	".............#.....##.........#...#...............................................................................................",
	".................................................................................#.............#.....#...#..#.....................",
	".................................#......#.#...#.......#......................................#......#..........................#..",
	"#..#....................................#..#......................................................................................",
	"...#......#.......................#..........................#...........#.#................#.#.......#...........................",
	"#....#...............#..........................#....................#....#................................#............##........",
	"...........#........................#..........#......................................#...................................#.....#.",
	"......#................................#................#.............................#.........#...#........#....................",
	".............#................#...........#.#...........#.......................................#...........#.#...................",
	".....#......................#..........................#.................................#.............#............#....#........",
	".....#.........#.#..............................................................................#...................##............",
	"......#.....#..#.............#...............#......#..#.#.................................#...............#...#..................",
	"....................#.......................#...................#........................#...................#..#.....##..........",
	"...............#...........#..#...........................................#........................#....#.........................",
	".....................................#...................#.#............#......................................#..................",
	".....#......................#.#....................#.............................................#...........................#..#.",
	"..........#........#................#.........#...........................................#....#..#...........#..####..........#..",
}

func TestExitLab(t *testing.T) {

	t.Run("sample map exit", func(t *testing.T) {
		ExitLab(sampleLab)
	})
	t.Run("full map exit", func(t *testing.T) {
		ExitLab(fullLab)
	})
	t.Run("sample map infinite loops", func(t *testing.T) {
		ExitLabWithObstruction(sampleLab)
	})
	t.Run("full map infinite loops", func(t *testing.T) {
		ExitLab(sampleLab)
	})

}