package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DcWire/path-finder-go/router"
)

func main() {
	// var maze = [][]int{
	// 	{0, 0, 0, 0, 0},
	// 	{0, -1, -1, 0, 0},
	// }

	// var vis = make([][]bool, len(maze))
	// for i, _ := range vis {
	// 	vis[i] = make([]bool, len(maze[0]))
	// }

	// var start = []int{0, 0}
	// var end = []int{len(maze) - 1, len(maze[0]) - 1}
	// var isPath bool = controller.Bfs(maze, start[0], start[1], end, vis)
	// fmt.Println(isPath)

	router := router.Router()

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Server is running on port 8000")

}
