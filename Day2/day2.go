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
		fmt.Println("IndividualGame")
		fmt.Println(IndividualGame)
		AllGames = append(AllGames, IndividualGame)
		fmt.Println("AllGames")
		fmt.Println(AllGames)
	}

	//for i, _ := range k2 {
	//	game := Game{}
	//	works := true
	//	splitSubstrings := strings.Split(strings.Join(k2[i], " "), ";")
	//	for _, o := range splitSubstrings {
	//		word := strings.Split(o, " ")
	//		fmt.Println("word")
	//		fmt.Println(word)
	//		for x, j := range word {
	//			if j == "blue" {
	//				valueOfBlue, _ := strconv.Atoi(word[x-1])
	//				if valueOfBlue > 14 {
	//					works = false
	//					fmt.Println(i)
	//					fmt.Println("changed to false")
	//
	//				}
	//			}
	//			if j == "red" {
	//				valueOfRed, _ := strconv.Atoi(word[x-1])
	//				if valueOfRed > 12 {
	//					works = false
	//					fmt.Println(i)
	//					fmt.Println("changed to false")
	//
	//				}
	//			}
	//			if j == "green" {
	//				valueOfGreen, _ := strconv.Atoi(word[x-1])
	//				if valueOfGreen > 13 {
	//					works = false
	//					fmt.Println(i)
	//					fmt.Println("changed to false")
	//				}
	//			}
	//		}
	//
	//	}
	//	if works == true {
	//		total += i
	//	}
	//	row = append(row, game)
	//}
	return total
}

func Run() int {
	gameData := CreateData()
	fmt.Println(gameData)

	total := 0
	return total
}
