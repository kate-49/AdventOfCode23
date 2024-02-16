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
	fmt.Println("gd")
	fmt.Println(gameData[0])
	fmt.Println(gameData[1])

	for _, row := range gameData {
		for _, eachNumber := range row.perRowCoordinates {
			isValid := CheckCoordinate(eachNumber, gameData)
			if isValid {
				fmt.Println("valid")
				fmt.Println(isValid)
			}
		}
	}
	return partNumbers
}

func CheckCoordinate(coordinates []int, gameData []RowElement) bool {
	fmt.Println("coord")
	fmt.Println(coordinates)
	startCoordinate := []int{coordinates[1], coordinates[2]}
	endCoordinate := []int{coordinates[3], coordinates[4]}
	fmt.Println("start")
	fmt.Println(startCoordinate)
	fmt.Println("end")
	fmt.Println(endCoordinate)

	////check element to the left if element is not 0
	//
	//if startCoordinate[0] != 0 {
	//	leftStartElement := startCoordinate[0] - 1
	//	fmt.Println("left")
	//	fmt.Println(leftStartElement)
	//	fmt.Println(startCoordinate[1])
	//	fmt.Println(gameData[1].stringRow)
	//	if (gameData[leftStartElement].stringRow[startCoordinate[1]]) == "*" {
	//		return true
	//	}
	//}

	////check we're not on row 0
	//if startCoordinate[1] != 0 {
	//	//check element above first element
	//
	//	aboveStartElement := startCoordinate[1] - 1
	//	fmt.Println("above start")
	//	fmt.Println(startCoordinate[0])
	//	fmt.Println(aboveStartElement)
	//	fmt.Println(gameData[startCoordinate[0]].stringRow[aboveStartElement])
	//	if (gameData[startCoordinate[0]].stringRow[aboveStartElement]) == "*" {
	//		return true
	//	}
	//	//check element above middle element
	//	//if (gameData[startCoordinate[0]].stringRow[aboveStartElement]) == "*" {
	//	//	return true
	//	//}
	//	//check element above last element
	//	aboveEndElement := endCoordinate[1] - 1
	//	fmt.Println("above end")
	//	fmt.Println(endCoordinate[0])
	//	fmt.Println(aboveEndElement)
	//	fmt.Println(gameData[endCoordinate[0]].stringRow[aboveEndElement])
	//	if (gameData[endCoordinate[0]].stringRow[aboveEndElement]) == "*" {
	//		return true
	//	}
	//
	//}

	//check we're not on final row
	if startCoordinate[1] < len(gameData)-1 {
		//check element below start element
		belowStartElement := startCoordinate[1] + 1
		if (gameData[belowStartElement].stringRow[startCoordinate[0]]) == "*" {
			return true
		}

		if (coordinates[4] - coordinates[2]) > 1 {
			//check element above middle element
			midCoordinate := startCoordinate[0] + 1
			if (gameData[belowStartElement].stringRow[midCoordinate]) == "*" {
				return true
			}
		}

		//check element below last element
		belowEndElement := endCoordinate[0] + 1
		if (gameData[belowEndElement].stringRow[endCoordinate[1]]) == "*" {
			return true
		}
	}

	////check we're not on final element in array
	////check element to the right of last element
	//if startCoordinate[0] < 9 {
	//	rightEndElement := startCoordinate[0] + 1
	//	fmt.Println("right")
	//	fmt.Println(endCoordinate[0])
	//	fmt.Println(rightEndElement)
	//	fmt.Println(gameData[endCoordinate[0]].stringRow[rightEndElement])
	//	if (gameData[endCoordinate[0]].stringRow[rightEndElement]) == "*" {
	//		return true
	//	}
	//}

	return false
}
