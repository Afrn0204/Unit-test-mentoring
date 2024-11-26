package main

import "fmt"

func luasBelahKetupat(diagonal1, diagonal2 float64) float64 {
	luas := (diagonal1 * diagonal2) / 2
	return luas
}

func main() {
	fmt.Println(luasBelahKetupat(10, 8))
}
