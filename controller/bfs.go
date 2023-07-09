package controller

func isValid(vis [][]bool, row int, col int, len1 int, len2 int) bool {
	if row >= len1 || row < 0 || col >= len2 || col < 0 {
		return false
	} else if vis[row][col] {
		return false
	}
	return true
}
func Bfs(maze [][]int, i int, j int, end []int, vis [][]bool) bool {
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
