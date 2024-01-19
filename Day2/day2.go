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

func CreateData() [][]string {
	var data [][]string
	file, _ := os.Open("day2pt1input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(scanner.Text(), "Game")
		data = append(data, entries)
	}

	_ = file.Close()

	var kate [][]string
	var k2 [][]string

	for i, _ := range data {
		subgame := []string{}
		entries := strings.Split(strings.Join(data[i], ","), " ")

		for w, v := range entries {
			viableSubGame := checkIfSubGameIsViable(clearString(v), w, entries)
			if !viableSubGame {
				continue
			}
			if w > 1 {
				subgame = append(subgame, clearString(v), entries[w+1])
			}
			endOfSubGame, _ := regexp.MatchString(";", v)
			if endOfSubGame == true {
				kate = append(kate, subgame)
				i++
				subgame = []string{}
			}
		}
		kate = append(kate, subgame)
		k2 = append(k2, kate[0])
		kate = [][]string{}
	}

	for i, _ := range k2 {
		joinedSubstrings := strings.Join(k2[i], " ")
		fmt.Println("joinedSubstrings")
		fmt.Println(joinedSubstrings)
		splitSubstrings := strings.Split(joinedSubstrings, ";")
		fmt.Println("splitSubstrings")
		fmt.Println(splitSubstrings[0])
	}
	return k2

	//for i, _ := range kate {
	//	game := Game{}
	//	for l, v := range kate[i] {
	//		fmt.Println("kate")
	//		fmt.Println(i)
	//		fmt.Println(kate[i])
	//
	//		if clearString(v) == "blue" {
	//			fmt.Println(kate[i][l-1])
	//			numOfBlue, _ := strconv.Atoi(kate[i][l-1])
	//			game.Blue = numOfBlue
	//		}
	//		if clearString(v) == "red" {
	//			numOfRed, _ := strconv.Atoi(kate[i][l-1])
	//			game.Red = numOfRed
	//		}
	//		if clearString(v) == "green" {
	//			numOfGreen, _ := strconv.Atoi(kate[i][l-1])
	//			game.Green = numOfGreen
	//		}
	//	}
	//	row = append(row, game)
	//}

	//return row

}

func Run() int {
	gameData := CreateData()
	fmt.Println(gameData)

	total := 0

	//for i := 0; i < len(gameData); i++ {
	//	gameIsViable := true
	//	//for l, v := range gameData {
	//	//	subgameIsViable := CheckIfGameIsPossible(cubesWeHave, gameData[i])
	//	//	if subgameIsViable == false {
	//	//		gameIsViable = false
	//	//	}
	//	//}
	//	//if gameIsViable == true {
	//	//	total += gameData[i].Id
	//	//}
	//}
	return total
}

func checkIfSubGameIsViable(input string, substringLocation int, entries []string) bool {
	cubesWeHave := Game{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	if clearString(input) == "blue" {
		numOfBlue, _ := strconv.Atoi(entries[substringLocation-1])
		if cubesWeHave.Blue > numOfBlue {
			return false
		}
	}
	if clearString(input) == "red" {
		numOfRed, _ := strconv.Atoi(entries[substringLocation-1])
		if cubesWeHave.Red > numOfRed {
			return false
		}
	}
	if clearString(input) == "green" {
		numOfGreen, _ := strconv.Atoi(entries[substringLocation-1])
		if cubesWeHave.Green > numOfGreen {
			return false
		}
	}
	return true
}
