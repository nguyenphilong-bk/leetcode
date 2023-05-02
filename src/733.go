package src

import "fmt"

func colorize(image [][]int, sr, sc, color, rLen, cLen int) [][]int {
    fmt.Println(sr, sc, rLen, cLen)
    if sr >= rLen || sr < 0 || sc >= cLen || sc < 0 {
        return image
    }
    
    if sr-1 >= 0 && image[sr][sc] == image[sr-1][sc] {
        image = colorize(image, sr-1, sc, color, rLen, cLen)
    }
    if sr + 1 < rLen && image[sr][sc] == image[sr+1][sc] {
        image = colorize(image, sr+1, sc, color, rLen, cLen)
    }
    if sc - 1 >= 0 && image[sr][sc] == image[sr][sc-1] {
        image = colorize(image, sr, sc-1, color, rLen, cLen)
    }

    // if sc + 1 < cLen && image[sr][sc] == image[sr][sc+1] {
    //     fmt.Println("coming")
    //     image = colorize(image, sr, sc+1, color, rLen, cLen)
    // }

    image[sr][sc] = color
    return image
}

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
    rLen := len(image)
    cLen := len(image[0])

    return colorize(image, sr, sc, color, rLen, cLen)    
}