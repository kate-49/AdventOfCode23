package Day_2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Red   []int
	Green []int
	Blue  []int
}

func clearString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func CreateData() int {
	var data [][]string
	file, _ := os.Open("day2pt1input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(scanner.Text(), "Game")
		data = append(data, entries)
	}

	_ = file.Close()

	total := 0

	var AllGames []Game

	for i, _ := range data {
		var IndividualGame Game

		entries := strings.Split(data[i][1], " ")
		gameId, _ := strconv.Atoi(clearString(entries[1]))
		IndividualGame.Id = gameId
		for w, v := range entries {
			containsGreen, _ := regexp.MatchString("green", v)
			containsBlue, _ := regexp.MatchString("blue", v)
			containsRed, _ := regexp.MatchString("red", v)

			if containsGreen {
				green, _ := strconv.Atoi(entries[w-1])
				IndividualGame.Green = append(IndividualGame.Green, green)
			}
			if containsBlue {
				blue, _ := strconv.Atoi(entries[w-1])
				IndividualGame.Blue = append(IndividualGame.Blue, blue)
			}
			if containsRed {
				red, _ := strconv.Atoi(entries[w-1])
				IndividualGame.Red = append(IndividualGame.Red, red)
			}
		}
		AllGames = append(AllGames, IndividualGame)
	}

	for _, g := range AllGames {
		works := true
		for _, subgameValue := range g.Blue {
			if subgameValue > 14 {
				works = false
			}
		}
		for _, subgameValue := range g.Red {
			if subgameValue > 12 {
				works = false
			}
		}
		for _, subgameValue := range g.Green {
			if subgameValue > 13 {
				works = false
			}
		}
		if works == true {
			total += g.Id
		}
	}
	return total
}

func Run() int {
	gameData := CreateData()
	return gameData
}
