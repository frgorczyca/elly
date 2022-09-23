package main

import (
	"fmt"
)

func main() {

	var nr = CalculateElo(1613, []MatchResult{{1609, 0}, {1477, 0.5}, {1388, 1}, {1586, 1}, {1720, 0}})

	fmt.Printf("this is new Elo %d", nr)
}
