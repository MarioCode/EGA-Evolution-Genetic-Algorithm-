package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	count_things   int = 15
	count_pop      int = 100
	len_population int = 15
	max_weight     int = 80
)

var (
	population     = make([][]int, count_pop)
	tmp_population = make([][]int, count_pop)
	things         = make([]int, count_things)
	price          = make([]int, count_things)
	cost_pop       = make([]int, count_pop)
	weight_pop     = make([]int, count_pop)
)

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Filling the the bag (on assignment)
func init_bag() {
	things[0], price[0] = 2, 21
	things[1], price[1] = 26, 19
	things[2], price[2] = 23, 27
	things[3], price[3] = 6, 3
	things[4], price[4] = 19, 24
	things[5], price[5] = 9, 30
	things[6], price[6] = 3, 6
	things[7], price[7] = 20, 13
	things[8], price[8] = 11, 2
	things[9], price[9] = 1, 21
	things[10], price[10] = 17, 26
	things[11], price[11] = 22, 26
	things[12], price[12] = 7, 24
	things[13], price[13] = 20, 1
	things[14], price[14] = 11, 20
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Creating the initial generation. Random or stochastic method
func generation(generation string) {

	// -r - random filling
	// -r - stochastic filling
	if generation == "-r" {
		for i := 0; i < count_pop; i++ {
			for {
				num_thing := rand.Intn(count_things)
				if population[i][num_thing] == 1 {
					continue
				} else if !is_full(population[i], num_thing) {
					population[i][num_thing] = 1
				} else {
					break
				}
			}
		}
	} else if generation == "-s" {
		for i := 0; i < count_pop; i++ {
			stochastic(population[i])
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Stochastic generation
func stochastic(object []int) {
	for i := 0; i < len(object); i++ {
		object[i] = rand.Intn(2)

		if (object[i] == 1) && (s_weight(object) > max_weight) {
			object[i] = 0
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Checking bags overload
func is_full(object []int, thing int) bool {

	var sum int
	tmp := make([]int, len_population)
	copy(tmp, object)
	tmp[thing] = 1

	for i := 0; i < len(tmp); i++ {
		sum += tmp[i] * things[i]
	}

	if sum > max_weight {
		return true
	}
	return false
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Bubble sorting (by value)
func cost_bubble_sort() {

	length := len(cost_pop) - 1
	for i := 0; i < len(cost_pop); i++ {
		for j := 0; j < length; j++ {
			if cost_pop[j] < cost_pop[j+1] {
				population[j], population[j+1] = population[j+1], population[j]
				cost_pop[j], cost_pop[j+1] = cost_pop[j+1], cost_pop[j]
				weight_pop[j], weight_pop[j+1] = weight_pop[j+1], weight_pop[j]
			}
		}
		length--
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Each population is calculated the weight and cost
func sum_weight_cost(object []int) (int, int) {
	var (
		sum_weight int
		sum_cost   int
	)

	for i := 0; i < len(object); i++ {
		sum_weight += object[i] * things[i]
		sum_cost += object[i] * price[i]
	}
	return sum_weight, sum_cost
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// The summation of the object by value (adaptability)
func s_cost(object []int) int {
	var sum_cost int

	for i := 0; i < len(object); i++ {
		sum_cost += object[i] * price[i]
	}
	return sum_cost
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// The summation of the object by weight
func s_weight(object []int) int {
	var sum_weight int

	for i := 0; i < len(object); i++ {
		sum_weight += object[i] * things[i]
	}
	return sum_weight
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Synchronization of populations with weight and value pack
func synchronization(object [][]int) {
	for i := range object {
		weight_pop[i], cost_pop[i] = sum_weight_cost(object[i])
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Displaying the results of intermediate calculations.
func Print(object [][]int) {

	synchronization(object)
	fmt.Println("---------------------")
	fmt.Println(things, " - Weight")
	fmt.Println(price, " - Cost")
	fmt.Println("---------------------")
	for i, s := range object {
		fmt.Printf("%v [%d - %d]\n", s, weight_pop[i], cost_pop[i])
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Selection Type - Tournament
// Select 2 random individuals, compares the value.
func tourney() {

	for i := 0; i < count_pop; i++ {
		pop_1 := rand.Intn(len_population)
		pop_2 := rand.Intn(len_population)

		sum_1, _ := sum_weight_cost(population[pop_1])
		sum_2, _ := sum_weight_cost(population[pop_2])

		if sum_1 >= sum_2 {
			copy(tmp_population[i], population[pop_1])
		} else {
			copy(tmp_population[i], population[pop_2])
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Selection Type - Roulette
func roulette() {
	var all_s float64
	var roulette_segment = make([]float64, count_pop)

	synchronization(population)
	cost_bubble_sort()

	for i := 0; i < count_pop; i++ {
		cur_sum := s_cost(population[i])
		full_sum := 0
		for j := 0; j < count_pop; j++ {
			if j == i {
				continue
			}
			full_sum += s_cost(population[j])
		}
		roulette_segment[i] = float64(cur_sum) / float64(full_sum)
	}

	for i := 0; i < count_pop; i++ {
		all_s += roulette_segment[i]
	}

	for i := 0; i < count_pop; i++ {
		roulette_segment[i] = roulette_segment[i] / all_s * 100
		//fmt.Printf("%d - %.4f\n", i, roulette_segment[i])
	}

	for i := 0; i < count_pop; i++ {
		diff := (rand.Float64() * (roulette_segment[0] - roulette_segment[count_pop-1])) + roulette_segment[count_pop-1]

		for j := count_pop - 1; j >= 0; j-- {
			if roulette_segment[j] < diff {
				continue
			} else {
				copy(tmp_population[i], population[j])
				break
			}
		}
		//fmt.Println(diff)
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Crossing
// -Single-point crossover
// Generate point gap, and to the breaking point two individuals swap content
//*******************************************//
// -Two-point crossover.
// In the main() generated two random points that make up the segment that will change when crossed.
// The beginning and the end is always inherited from the first parents, and the interval (a set of bits - segment), exchanges between parents.
func crossing(start, end, par1, par2 int, typeCross string) {

	if typeCross == "-two" {
		for j := start; j < end; j++ {
			tmp_population[par1][j], tmp_population[par2][j] = tmp_population[par2][j], tmp_population[par1][j]
		}

	} else if typeCross == "-one" {
		swap := rand.Intn(len_population/2) + len_population/4

		for j := 0; j < swap; j++ {
			tmp_population[par1][j], tmp_population[par2][j] = tmp_population[par2][j], tmp_population[par1][j]
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// The mutation (gene inversion) of the new individual in 5 random positions with a probability of 0.3
func mutation(object []int) {

	for j := 0; j < 5; j++ {
		percent_mut := rand.Intn(10)
		count_pos := rand.Intn(len_population)

		if percent_mut <= 3 {
			if object[count_pos] == 1 {
				object[count_pos] = 0
			} else {
				object[count_pos] = 1
			}
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// Finding the maximum, minimum, difference and the dispersion (sigma).
func math_stat() (int, int, float64, float64) {
	max := cost_pop[0]
	min := cost_pop[count_pop-1]

	var medium float64
	var sigma float64
	var disper float64

	for i := 0; i < count_pop; i++ {
		medium += float64(cost_pop[i])
	}
	medium = medium / float64(count_pop)

	for i := 0; i < count_pop; i++ {
		disper += (float64(cost_pop[i]) - medium) * (float64(cost_pop[i]) - medium)
	}
	sigma = disper / float64(count_pop)

	return min, max, medium, math.Sqrt(sigma)
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
// The main function that performs the functions of the genetic algorithm
func run_GA(start, end, num_finish int, typeCross, selection string) {

	var par1, par2, count int

	for count != num_finish {

		// -r Roulette
		// -t Tournament
		if selection == "-t" {
			tourney()
		} else if selection == "-r" {
			roulette()
		}

		for i := 0; i < count_pop/4; i++ {
			for {
				par1 = rand.Intn(count_pop)
				par2 = rand.Intn(count_pop)

				if par1 != par2 {
					break
				}
			}

			old1 := make([]int, len_population)
			old2 := make([]int, len_population)
			copy(old1, tmp_population[par1])
			copy(old2, tmp_population[par2])

			//Crossing
			crossing(start, end, par1, par2, typeCross)

			// Mutation
			mutation(tmp_population[par1])
			mutation(tmp_population[par2])

			// If the total weight exceeds, the generation of "fall off" and generated new.
			if s_weight(tmp_population[par1]) > max_weight || s_weight(tmp_population[par2]) > max_weight {
				copy(tmp_population[par1], old1)
				copy(tmp_population[par2], old2)
				//continue
			}
		}
		//synchronization(population)
		//roulette()

		for i := 0; i < count_pop; i++ {
			copy(population[i], tmp_population[i])
		}

		synchronization(population)
		cost_bubble_sort()
		Print(population)
		count += 1
	}
	min, max, medium, sig := math_stat()
	fmt.Printf("Min: %d\nMax: %d\nMed: %f\nSig: %f", min, max, medium, sig)
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
//main()
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var typeCross, generations, selections string
	var start, end, num_finish int

	if len(os.Args) > 3 {

		generations = os.Args[1]
		typeCross = os.Args[2]
		selections = os.Args[3]
	} else {
		fmt.Println("Please, enter the parameters (-r/-s -one/-two -t/-r")
		os.Exit(2)
	}

	for i := 0; i < count_pop; i++ {
		population[i] = make([]int, len_population)
		tmp_population[i] = make([]int, len_population)
	}

	//Initialization bags
	init_bag()

	//Creating the initial generation.
	generation(generations)

	// Generation section for two-point crossover
	start = rand.Intn(len_population/2) + 1
	end = rand.Intn(len_population/2) + len_population/2 - 1
	num_finish = 10

	Print(population)
	cost_bubble_sort()
	Print(population)

	run_GA(start, end, num_finish, typeCross, selections)
}
