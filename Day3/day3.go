package Day_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RowElement struct {
	perRowCoordinates [][]int
	numbers           []int
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
	rowNumber := 0

	for scanner.Scan() {
		entries := scanner.Text()

		numbers := getWholeNumbersFromRowInput(entries)
		wholeRowAsStringArray := strings.Split(entries, "")
		//change this so get the first whole number then loop over the string array to see where it matches, check the following elements match and if so add coordinates

		element := RowElement{
			numbers: numbers,
		}
		var coord []int

		for _, num := range numbers {
			numberAsIntArray := strings.Split(strconv.Itoa(num), "")
			for k, _ := range wholeRowAsStringArray {
				intLength := len(numberAsIntArray)
				if numberAsIntArray[0] == wholeRowAsStringArray[k] {
					if intLength > 1 {
						if numberAsIntArray[1] == wholeRowAsStringArray[k+1] {
							if intLength > 2 {
								if numberAsIntArray[2] == wholeRowAsStringArray[k+2] {
									coord = []int{num, rowNumber, k, rowNumber, k + 2}
								}
							} else {
								coord = []int{num, rowNumber, k, rowNumber, k + 1}
							}
						}
					}
				}
				if len(coord) > 1 {
					//check if contains element before appending
					alreadySaved := checkForDuplicateElements(element.perRowCoordinates, coord)
					if !alreadySaved {
						element.perRowCoordinates = append(element.perRowCoordinates, coord)
					}
				}

			}
		}
		if len(element.numbers) >= 1 {
			data = append(data, element)
		}
		rowNumber++
	}

	_ = file.Close()
	fmt.Println("close file")
	return data
}

func checkForDuplicateElements(existingCoordinatesForRow [][]int, coord []int) bool {
	for _, el := range existingCoordinatesForRow {
		if (el[0] == coord[0]) && (el[1] == coord[1]) && (el[2] == coord[2]) {
			return true
		}
	}
	return false
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

	for p, row := range gameData {
		fmt.Println("p")
		fmt.Println(p)
		fmt.Println("row")
		fmt.Println(row)
	}
	//	//for k := 0; k < len(row.numbersWithYIndex); k++ {
	//	//	fmt.Println("el")
	//	//	fmt.Println(row.numbersWithYIndex[k])
	//	//	//if el[0] == 1 {
	//	//	//	matchingRows = append(matchingRows, row)
	//	//	//}
	//	//}
	//}
	////get number location
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
	//if x > 0 && x+1 < len(gameData[x]) {
	//	if (gameData[x+1][y] != ".") && (!contains(validLetters, gameData[x+1][y])) {
	//		fmt.Println("true")
	//		fmt.Println(gameData[x][y])
	//		partNumber = true
	//	}
	//}
	//if x-1 > 0 {
	//	if (gameData[x-1][y] != ".") && (!contains(validLetters, gameData[x-1][y])) {
	//		fmt.Println("true")
	//		fmt.Println(gameData[x][y])
	//		partNumber = true
	//	}
	//}
	fmt.Println("----")
	return partNumber
}
