package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Population struct {
	len_pop, best_result int
	population, weight   []int
	neighbors            [][]int
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Adaptability of population
func (p *Population) sum_weight(pop int) int {
	var sum int
	for i := range p.weight {
		sum += (p.neighbors[pop][i] * p.weight[i])
	}

	return sum
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Run Hill Climbing
func (p *Population) go_climbing() {

	max_pos := 0
	max := p.sum_weight(max_pos)
	count := 0

	for {
		for i := 0; i < p.len_pop; i++ {

			//Neighbors generation
			copy(p.neighbors[i], p.population)

			if p.neighbors[i][i] == 0 {
				p.neighbors[i][i] = 1
			} else {
				p.neighbors[i][i] = 0
			}

			// Finding the best populations
			cur := p.sum_weight(i)

			if cur > max {
				max_pos = i
				max = cur
			}
		}

		// If the best of all the neighbors not found - return.
		if max <= p.best_result {
			fmt.Printf("Optimal solution is found.\nSteps needed: %d\n\n", count)
			return
		}

		copy(p.population, p.neighbors[max_pos])
		p.Print(max_pos)
		p.best_result = max
		count += 1
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Print results
func (p *Population) Print(max int) {
	for i := 0; i < p.len_pop; i++ {
		if i == max {
			fmt.Printf("%v - [%d] - Best result\n", p.neighbors[i], p.sum_weight(i))
		} else {
			fmt.Printf("%v - [%d]\n", p.neighbors[i], p.sum_weight(i))
		}
	}
	fmt.Printf("\n")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Population size: default is 10
	// Other size via parameter Args
	count_pop := 10
	if len(os.Args) > 0 && len(os.Args) <= 2 {
		tmp, _ := strconv.Atoi(os.Args[1])
		if tmp >= 5 {
			count_pop = tmp
		}
	}

	pop := Population{len_pop: count_pop, best_result: 0}
	pop.population = make([]int, pop.len_pop)
	pop.weight = make([]int, pop.len_pop)
	pop.neighbors = make([][]int, pop.len_pop)

	for i := range pop.population {
		pop.population[i] = rand.Intn(2)
		pop.weight[i] = rand.Intn(5) + 1
		pop.neighbors[i] = make([]int, pop.len_pop)
	}

	fmt.Printf("\n%v - Initial population\n%v - Weight of population\n\n", pop.population, pop.weight)

	pop.go_climbing()
}
