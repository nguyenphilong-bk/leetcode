package main

import (
	"fmt"
	"leetcode/src"
)

func main() {
	// image := [[1,1,1],[1,1,0],[1,0,1]]
	image := [][]byte{
		{'1','1','1'},
		{'1','1','0'},
		{'1','0','1'},
	}	
	// src.FloodFill(image, 1, 1, 2)
	result := src.NumIslands(image)
	fmt.Print(result)
}