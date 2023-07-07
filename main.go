package main

import "fmt"

func dfs(maze [][]int, i int, j int, end []int, visPtr *[][]bool) bool {

	vis := *visPtr
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
	if dfs(maze, i+1, j, end, &vis) || dfs(maze, i-1, j, end, &vis) || dfs(maze, i, j+1, end, &vis) || dfs(maze, i, j-1, end, &vis) {
		return true
	}

	vis[i][j] = false

	return false
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
	var isPath bool = dfs(maze, start[0], start[1], end, &vis)
	fmt.Println(isPath)

}
