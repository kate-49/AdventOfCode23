package Day_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RowElement struct {
	xCoord            int
	wholeNumber       int
	rowAsStringArray  []string
	numbersWithYIndex [][]int
}

func GetNumberOfDigitsPerRow(numbers []int) int {
	kate := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ""), "[]")
	return len(kate)
}

func CreateData() []RowElement {
	var data []RowElement
	file, _ := os.Open("day3input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	j := 0

	for scanner.Scan() {
		entries := scanner.Text()
		numbers := getWholeNumbersFromRowInput(entries)
		numbersOfIndividualDigitsPerRow := GetNumberOfDigitsPerRow(numbers)
		//this needs to be for the number of indiviual digits in numbers
		fmt.Println("elements in row")
		fmt.Println(numbersOfIndividualDigitsPerRow)

		for i := 0; i < numbersOfIndividualDigitsPerRow; i++ {
			fmt.Println("i")
			fmt.Print(i)
			element := RowElement{}
			element.xCoord = j
			element.wholeNumber = numbers[j]
			element.rowAsStringArray = strings.Split(entries, "")
			for d, h := range element.rowAsStringArray {
				hAsInt, _ := strconv.Atoi(h)
				if hAsInt != 0 {
					element.numbersWithYIndex = append(element.numbersWithYIndex, []int{hAsInt, d, i})
				}
			}
			data = append(data, element)
		}
		j++
	}

	_ = file.Close()
	return data
}

func getWholeNumbersFromRowInput(input string) []int {
	finalElements := []int{}

	intputAsArray := strings.Split(input, ".")
	for _, el := range intputAsArray {
		elAsInt, _ := strconv.Atoi(el)
		if elAsInt != 0 {
			finalElements = append(finalElements, elAsInt)
		}
	}
	return finalElements
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Run() []int {
	partNumbers := []int{}

	gameData := CreateData()

	fmt.Println("game data")
	fmt.Println(gameData)
	for _, row := range gameData {
		fmt.Println(row.wholeNumber)
		fmt.Println(row.numbersWithYIndex)
	}
	//get number location
	//check row above and below
	//check current row
	//k2 := []Kate{}
	//kate := Kate{}
	//
	//for elementIndex, element := range row {
	//	//group numbers together
	//	//joined := strings.Join(row, "")
	//	//numbers := strings.Split(joined, ".")
	//	//
	//	//for _, NumElement := range numbers {
	//	//	fmt.Println("NumElement")
	//	//	fmt.Println(NumElement)
	//	//	if contains(validLetters, NumElement) {
	//	//		kate.wholeNumber = NumElement
	//	//		fmt.Println("elementIndex")
	//	//		fmt.Println(elementIndex)
	//	//		fmt.Println("el")
	//	//		fmt.Println(element)
	//	//	}
	//	//}
	//
	//	//if contains(validLetters, element) {
	//	//	isPartNumber := CheckCoordinate([]int{index, elementIndex}, gameData, validLetters)
	//	//	if isPartNumber {
	//	//		intValue := getRealValue(element, numbers)
	//	//		partNumbers = append(partNumbers, intValue)
	//	//	}
	//	//}
	//
	//}
	//}
	return partNumbers
}

func CheckCoordinate(coordinates []int, gameData [][]string, validLetters []string) bool {
	partNumber := false
	x := coordinates[0]
	y := coordinates[1]
	fmt.Println(gameData[x][y])
	if x > 0 && x+1 < len(gameData[x]) {
		if (gameData[x+1][y] != ".") && (!contains(validLetters, gameData[x+1][y])) {
			fmt.Println("true")
			fmt.Println(gameData[x][y])
			partNumber = true
		}
	}
	if x-1 > 0 {
		if (gameData[x-1][y] != ".") && (!contains(validLetters, gameData[x-1][y])) {
			fmt.Println("true")
			fmt.Println(gameData[x][y])
			partNumber = true
		}
	}
	fmt.Println("----")
	return partNumber
}
