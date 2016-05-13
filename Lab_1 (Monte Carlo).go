package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	count_pop  = 100
	len_pop    = 5
	population = make([][]int, count_pop)
	weight     = make([]int, len_pop)
)

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Adaptability of population
func find_sum(pos int) int {

	var sum int
	for i := 0; i < len_pop; i++ {
		sum += population[pos][i] * weight[i]
	}

	return sum
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// main() :)
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < count_pop; i++ {
		population[i] = make([]int, len_pop)
		for j := 0; j < len_pop; j++ {
			population[i][j] = rand.Intn(2)
		}
	}
	for i := 0; i < len_pop; i++ {
		weight[i] = rand.Intn(10)
	}

	fmt.Println("Weight: ", weight, "\n")

	for i := 0; i < count_pop; i++ {
		fmt.Printf("[%d] - %v\n", i, population[i])
	}

	var best_pos, cur int
	fmt.Println("\nTop scores: \n")

	fmt.Println("Weight: ", weight, "\n")
	for i := 0; i < 5; i++ {
		pos := rand.Intn(count_pop)
		cur = find_sum(pos)

		if cur >= best_pos {
			best_pos = cur
			fmt.Printf("[%d] -\t%v - [%d] - Best in the current step (%d)\n", pos, population[pos], best_pos, best_pos)
		} else {
			fmt.Printf("[%d] -\t%v - [%d] - Best in the current step (%d)\n", cur, population[cur], find_sum(cur), best_pos)
		}
	}
}
