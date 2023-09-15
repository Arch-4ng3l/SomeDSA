package algorithms

type Point struct {
	X int
	Y int
}

var dir = []Point{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func walk(maze []string, wall string, cur Point, end Point, seen [][]bool, path []Point) ([]Point, bool) {

	if cur.X < 0 || cur.X >= len(maze[0]) || cur.Y < 0 || cur.Y >= len(maze) {
		return nil, false
	}

	if string(maze[cur.Y][cur.X]) == wall {
		return nil, false
	}

	if cur.X == end.X && cur.Y == end.Y {
		path = append(path, cur)
		return path, true
	}

	if seen[cur.Y][cur.X] {
		return nil, false
	}

	path = append(path, cur)

	seen[cur.Y][cur.X] = true

	for _, d := range dir {
		if curPath, b := walk(maze, wall, Point{X: cur.X + d.X, Y: cur.Y + d.Y}, end, seen, path); b {
			return curPath, true
		}
	}

	path = path[:len(path)-1]

	return nil, false
}

func MazeSolve(maze []string, wall string, start Point, end Point) ([]Point, []string) {
	seen := make([][]bool, len(maze))

	for i := range maze {
		seen[i] = make([]bool, len(maze[i]))
	}
	path := make([]Point, 0)

	path, _ = walk(maze, wall, start, end, seen, path)

	for i, p := range path {
		r := []rune(maze[p.Y])
		r[p.X] = rune(i + 64)
		maze[p.Y] = string(r)
	}

	return path, maze

}
