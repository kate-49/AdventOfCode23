package Day_2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Blue  int
	Red   int
	Green int
}

func clearString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func CreateData() []Game {
	var data [][]string
	var allData []Game
	file, _ := os.Open("day2pt1input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		data = append(data, entries)
	}

	_ = file.Close()

	for _, s := range data {
		game := Game{}
		//break into sub arrays based on the colons?
		for k, l := range s {
			if clearString(l) == "Game" {
				game.Id, _ = strconv.Atoi(clearString(s[k+1]))
			}
			if clearString(l) == "blue" {
				numOfBlue, _ := strconv.Atoi(s[k-1])
				game.Blue += numOfBlue
			}
			if clearString(l) == "red" {
				numOfRed, _ := strconv.Atoi(s[k-1])
				game.Red += numOfRed
			}
			if clearString(l) == "green" {
				numOfGreen, _ := strconv.Atoi(s[k-1])
				game.Green += numOfGreen
			}
		}
		allData = append(allData, game)
	}
	return allData
}

func Run() int {
	gameData := CreateData()
	fmt.Println(gameData)

	total := 0

	cubesWeHave := Game{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	for i := 0; i < len(gameData); i++ {
		total += CheckIfGameIsPossible(cubesWeHave, gameData[i])
	}
	return total
}

func CheckIfGameIsPossible(cubesWeHave Game, game Game) int {
	if cubesWeHave.Red >= game.Red && cubesWeHave.Green >= game.Green && cubesWeHave.Blue >= cubesWeHave.Blue {
		fmt.Println(game.Id)
		return game.Id
	}
	return 0
}
