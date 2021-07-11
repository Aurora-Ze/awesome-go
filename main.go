package main

import "fmt"

func main() {
	fmt.Println(movingCount(1, 3, 11))
}

var visited [][]bool
var result int
var direct [][]int
var M, N int

func movingCount(m int, n int, k int) int {
	visited = make([][]bool, m+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}
	fmt.Printf("%v", visited)

	direct = [][]int{{1, 0}, {0, 1}}
	result = 0
	M = m
	N = n
	dfs(0, 0, k, visited)

	return result
}

func dfs(x int, y int, k int, visited [][]bool) {
	fmt.Println(x, " ", y)
	if !inArea(x, y) || visited[x][y] {
		return
	}

	if x%10+x/10+y%10+y/10 <= k {
		result++
		visited[x][y] = true
		for i := 0; i < 2; i++ {
			dfs(x+direct[i][0], y+direct[i][1], k, visited)
		}
	}
}

func inArea(x, y int) bool {
	return x >= 0 && y >= 0 && x < M && y < N
}
