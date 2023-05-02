package src

import "fmt"

func dfs(grid [][]byte, r, c int, visited map[string]bool) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}

	key := fmt.Sprintf("%v-%v", r, c)
	if _, ok := visited[key]; ok {
		return
	}
	
	visited[key] = true
	if grid[r][c] == '1' {

		dfs(grid, r-1, c, visited)
		fmt.Println("first:", visited)
		dfs(grid, r+1, c, visited)
		fmt.Println("second:", visited)
		dfs(grid, r, c-1, visited)
		fmt.Println("third:", visited)
		dfs(grid, r, c+1, visited)
		fmt.Println("last:", visited)
	}
	fmt.Println("after:", visited)
}

func NumIslands(grid [][]byte) int {
	result := 0
	visited := make(map[string]bool, 0)

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			key := fmt.Sprintf("%v-%v", r, c)
			_, ok := visited[key]
			if grid[r][c] == '1' && !ok {
				dfs(grid, r, c, visited)
				result++
			}
		}
	}

	return result
}
