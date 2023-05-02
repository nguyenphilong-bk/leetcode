package src

import "fmt"

func colorize(image [][]int, sr, sc, color, rLen, cLen int, visited map[string]bool) [][]int {
	key := fmt.Sprintf("%v-%v", sr, sc)
	if _, ok := visited[key]; ok {
		return image
	} else {
		visited[key] = true
	}

	if sr >= rLen || sr < 0 || sc >= cLen || sc < 0 {
		return image
	}

	if sr-1 >= 0 && image[sr][sc] == image[sr-1][sc] {
		image = colorize(image, sr-1, sc, color, rLen, cLen, visited)
		fmt.Printf("image: %#v", image)
	}
	if sr+1 < rLen && image[sr][sc] == image[sr+1][sc] {
		image = colorize(image, sr+1, sc, color, rLen, cLen, visited)
	}
	if sc-1 >= 0 && image[sr][sc] == image[sr][sc-1] {
		image = colorize(image, sr, sc-1, color, rLen, cLen, visited)
	}

	if sc + 1 < cLen && image[sr][sc] == image[sr][sc+1] {
	    image = colorize(image, sr, sc+1, color, rLen, cLen, visited)
	}

	image[sr][sc] = color
	return image
}

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	rLen := len(image)
	cLen := len(image[0])
	visited := map[string]bool{}

	return colorize(image, sr, sc, color, rLen, cLen, visited)
}
