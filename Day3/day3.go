package Day_3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RowElement struct {
	perRowCoordinates [][]int
	numbers           []int
	stringRow         []string
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
			numbers:   numbers,
			stringRow: wholeRowAsStringArray,
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
								fmt.Println(num)
								if numberAsIntArray[2] == wholeRowAsStringArray[k+2] {
									fmt.Println("adding coord p2")
									coord = []int{num, k, rowNumber, k + 2, rowNumber}
									fmt.Println(coord)
								}
							} else {
								fmt.Println("adding coord p1")
								coord = []int{num, k, rowNumber, k + 1, rowNumber}
								fmt.Println(coord)

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
		data = append(data, element)

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

func checkIfContainsSymbol(input string) bool {
	validCharacters := []string{"*", "$", "#", "+"}

	for _, el := range validCharacters {
		if el == input {
			return true
		}
	}
	return false
}

func getWholeNumbersFromRowInput(input string) []int {
	finalElements := []int{}
	intputAsArray := strings.Split(input, ".")
	for _, el := range intputAsArray {
		elAsInt, _ := strconv.Atoi(clearString(el))
		if elAsInt != 0 {
			finalElements = append(finalElements, elAsInt)
		}
	}
	return finalElements
}

func clearString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func Run() int {
	gameData := CreateData()
	total := 0

	for _, row := range gameData {
		for _, eachNumber := range row.perRowCoordinates {
			isValid := CheckCoordinate(eachNumber, gameData)
			if isValid > 0 {
				fmt.Println("valid")
				fmt.Println(isValid)
				total += isValid
			}
		}
	}
	return total
}

func CheckCoordinate(coordinates []int, gameData []RowElement) int {
	startCoordinate := []int{coordinates[1], coordinates[2]}
	endCoordinate := []int{coordinates[3], coordinates[4]}

	//check element to the left if element is not 0
	if startCoordinate[0] > 0 {
		leftStartElement := startCoordinate[0] - 1
		if (checkIfContainsSymbol(gameData[startCoordinate[1]].stringRow[leftStartElement])) == true {
			return coordinates[0]
		}
	}

	//check element to the right of last element
	if startCoordinate[0] < 9 {
		rightEndElement := endCoordinate[0] + 1
		if (checkIfContainsSymbol(gameData[endCoordinate[1]].stringRow[rightEndElement])) == true {
			return coordinates[0]
		}
	}

	//check we're not on row 0
	if startCoordinate[1] != 0 {
		//check element above start element
		yCoordForRowAbove := startCoordinate[1] - 1
		if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[startCoordinate[0]])) == true {
			return coordinates[0]
		}

		if (coordinates[4] - coordinates[2]) > 1 {
			//check element above middle element
			midCoordinate := startCoordinate[0] + 1
			if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[midCoordinate])) == true {
				return coordinates[0]
			}
		}

		//check element above last element
		if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[endCoordinate[0]])) == true {
			return coordinates[0]
		}

	}

	//check we're not on final row
	if startCoordinate[1] < len(gameData)-1 {
		//check element below start element
		yCoordForRowBelow := startCoordinate[1] + 1
		if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[startCoordinate[0]])) == true {
			return coordinates[0]
		}

		if (coordinates[4] - coordinates[2]) > 1 {
			//check element below middle element
			midCoordinate := startCoordinate[0] + 1
			if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[midCoordinate])) == true {
				return coordinates[0]
			}
		}

		//check element below last element
		if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[endCoordinate[1]])) == true {
			return coordinates[0]
		}
	}

	return 0
}
