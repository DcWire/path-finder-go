package controller

func Dfs(maze [][]int, i int, j int, end []int, vis [][]bool) bool {

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
	if Dfs(maze, i+1, j, end, vis) || Dfs(maze, i-1, j, end, vis) || Dfs(maze, i, j+1, end, vis) || Dfs(maze, i, j-1, end, vis) {
		return true
	}

	vis[i][j] = false

	return false
}
