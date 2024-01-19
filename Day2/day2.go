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

func CreateData() int {
	var data [][]string
	var row []Game
	file, _ := os.Open("day2pt1input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(scanner.Text(), "Game")
		data = append(data, entries)
	}

	_ = file.Close()

	var kate [][]string
	var k2 [][]string
	total := 0

	for i, _ := range data {
		subgame := []string{}
		entries := strings.Split(strings.Join(data[i], ","), " ")

		for w, v := range entries {
			fmt.Println("V")
			fmt.Println(v)
			containsGreen, _ := regexp.MatchString("green", v)
			containsBlue, _ := regexp.MatchString("blue", v)
			containsRed, _ := regexp.MatchString("red", v)

			if containsGreen || containsBlue || containsRed {
				if w > 0 {
					if w < len(entries)-1 {
						fmt.Println("w")
						fmt.Println(entries[w+1])
						subgame = append(subgame, clearString(v), entries[w+1])
					}
				}
				endOfSubGame, _ := regexp.MatchString(";", v)
				if endOfSubGame == true {
					kate = append(kate, subgame)
					i++
					subgame = []string{}
				}
			}
		}
		kate = append(kate, subgame)
		fmt.Println("kate")
		fmt.Println(kate)
		k2 = append(k2, kate[0])
		kate = [][]string{}
	}

	for i, _ := range k2 {
		game := Game{}
		works := true
		splitSubstrings := strings.Split(strings.Join(k2[i], " "), ";")
		for _, o := range splitSubstrings {
			word := strings.Split(o, " ")
			fmt.Println("word")
			fmt.Println(word)
			for x, j := range word {
				if j == "blue" {
					valueOfBlue, _ := strconv.Atoi(word[x-1])
					if valueOfBlue > 14 {
						works = false
						fmt.Println(i)
						fmt.Println("changed to false")

					}
				}
				if j == "red" {
					valueOfRed, _ := strconv.Atoi(word[x-1])
					if valueOfRed > 12 {
						works = false
						fmt.Println(i)
						fmt.Println("changed to false")

					}
				}
				if j == "green" {
					valueOfGreen, _ := strconv.Atoi(word[x-1])
					if valueOfGreen > 13 {
						works = false
						fmt.Println(i)
						fmt.Println("changed to false")
					}
				}
			}

		}
		if works == true {
			total += i
		}
		row = append(row, game)
	}
	return total
}

func Run() int {
	gameData := CreateData()
	fmt.Println(gameData)

	total := 0
	return total
}
