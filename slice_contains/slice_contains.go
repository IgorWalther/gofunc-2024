package main

import (
	"fmt"
	"slices"
)

func SliceContainsV0(s []uint8, target uint8) bool {
	return slices.Contains(s, target)
}

func SliceContainsV1(s []uint8, target uint8) bool

func main() {
	fmt.Println(SliceContainsV1([]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, 2))
}
