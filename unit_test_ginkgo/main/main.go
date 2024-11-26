package main

import "fmt"

func LuasBelahKetupat(diagonal1, diagonal2 float64) float64 {
	luas := (diagonal1 * diagonal2) / 2
	return luas
}

func main() {
	fmt.Println(LuasBelahKetupat(10, 8))
}
