package Day_2

import (
	"sort"
)

func CalculateMinimumNeeded(AllGames []Game) []int {
	total := []int{}

	for _, g := range AllGames {

		//return max of each colour in each game
		sort.Ints(g.Blue)
		maximumBlue := g.Blue[len(g.Blue)-1]

		sort.Ints(g.Red)
		maximumRed := g.Red[len(g.Red)-1]

		sort.Ints(g.Green)
		maximumGreen := g.Green[len(g.Green)-1]

		gameTotal := maximumRed * maximumGreen * maximumBlue
		total = append(total, gameTotal)
	}
	return total
}

func SumTotalValue(powers []int) int {
	total := 0

	for _, p := range powers {
		total += p
	}
	return total
}

func RunDay2() int {
	gameData := CreateData()
	subgames := BreakIntoSubgames(gameData)
	arrayOfPowers := CalculateMinimumNeeded(subgames)
	return SumTotalValue(arrayOfPowers)
}
