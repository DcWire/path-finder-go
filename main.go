package main

import "fmt"

func dfs(maze [][]int, i int, j int, end []int, vis [][]bool) bool {

	if i == end[0] && j == end[1] {
		vis[i][j] = true
		return true
	}
	if i < 0 || j < 0 || i >= len(maze) || j >= len(maze[0]) {
		return false
	}

	if vis[i][j] {
		return false
	} else if maze[i][j] == -1 {
		return false
	}
	// fmt.Println("Row: ", i, "Col: ", j)
	vis[i][j] = true
	if dfs(maze, i+1, j, end, vis) || dfs(maze, i-1, j, end, vis) || dfs(maze, i, j+1, end, vis) || dfs(maze, i, j-1, end, vis) {
		return true
	}

	vis[i][j] = false

	return false
}

func isValid(vis [][]bool, row int, col int, len1 int, len2 int) bool {
	if row >= len1 || row < 0 || col >= len2 || col < 0 {
		return false
	} else if vis[row][col] {
		return false
	}
	return true
}
func bfs(maze [][]int, i int, j int, end []int, vis [][]bool) bool {
	dRow := []int{-1, 0, 1, 0}
	dCol := []int{0, 1, 0, -1}

	q := [][]int{}
	vis[i][j] = true
	q = append(q, []int{i, j})
	for len(q) > 0 {
		ele := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			r := ele[0] + dRow[i]
			c := ele[1] + dCol[i]
			if isValid(vis, r, c, len(maze), len(maze[0])) {
				q = append(q, []int{r, c})
				vis[r][c] = true
			}
		}
	}
	return vis[end[0]][end[1]]

}
func main() {
	var maze = [][]int{
		{0, 0, 0, 0, 0},
		{0, -1, -1, 0, 0},
	}

	var vis = make([][]bool, len(maze))
	for i, _ := range vis {
		vis[i] = make([]bool, len(maze[0]))
	}

	var start = []int{0, 0}
	var end = []int{len(maze) - 1, len(maze[0]) - 1}
	var isPath bool = bfs(maze, start[0], start[1], end, vis)
	fmt.Println(isPath)

}
