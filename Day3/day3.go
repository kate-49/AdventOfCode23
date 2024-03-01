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
	file, _ := os.Open("day3input2.txt")
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
			fmt.Println("num")
			fmt.Println(num)
			numberAsIntArray := strings.Split(strconv.Itoa(num), "")
			fmt.Println("numberAsIntArray")
			fmt.Println(numberAsIntArray)
			//for k, _ := range wholeRowAsStringArray {
			//runs through this per number in the numbers as int array
			for k := 0; k <= 139; k++ {
				fmt.Println("k")
				fmt.Println(k)
				intLength := len(numberAsIntArray)
				fmt.Println("after int length")
				//if this goes above 136 it times out but should go to 139?
				if numberAsIntArray[0] == wholeRowAsStringArray[k] {
					fmt.Println("num as int length")
					if intLength > 1 && k+1 <= 139 {
						//maybe debug here
						if numberAsIntArray[1] == wholeRowAsStringArray[k+1] {
							fmt.Println("num as int array 1")

							if intLength > 2 && k+2 <= 139 {
								fmt.Println("int length more than 2")

								fmt.Println(num)
								if numberAsIntArray[2] == wholeRowAsStringArray[k+2] {
									fmt.Println("adding coord p3")
									coord = []int{num, k, rowNumber, k + 2, rowNumber}
									fmt.Println(coord)
								}

							} else {
								fmt.Println("adding coord p2")
								coord = []int{num, k, rowNumber, k + 1, rowNumber}
								fmt.Println(coord)

							}
						}
						fmt.Println("other")
					} else if intLength == 1 {
						fmt.Println("adding coord p1")
						coord = []int{num, k, rowNumber, k, rowNumber}
						fmt.Println(coord)
					}
					if len(coord) > 1 {
						//check if contains element before appending
						alreadySaved := checkForDuplicateElements(element.perRowCoordinates, coord)
						fmt.Println("after already saved")
						if !alreadySaved {
							fmt.Println("not already saved")
							element.perRowCoordinates = append(element.perRowCoordinates, coord)
						} else {
							fmt.Println("continue")
							continue
						}
						//	add else continue?
					}
				}
				fmt.Println("end of loop")
				//}
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
		if (el[0] == coord[0]) && (el[1] == coord[1]) && (el[2] == coord[2]) && (el[3] == coord[3]) && (el[4] == coord[4]) {
			fmt.Println("return true")
			return true
		}
	}
	fmt.Println("return false")
	return false
}

func checkIfContainsSymbol(input string) bool {
	validCharacters := []string{"*", "$", "#", "+", "@", "/", "%", "=", "-", "-", "_", "&"}

	for _, el := range validCharacters {
		if el == input {
			return true
		}
	}
	return false
}

func getWholeNumbersFromRowInput(input string) []int {
	finalElements := []int{}
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}

	numericStr := reg.ReplaceAllString(input, ".")

	intputAsArray := strings.Split(numericStr, ".")
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
		//fmt.Println("coordinates per row")
		fmt.Println("total")
		fmt.Println(total)
		fmt.Println("new row------------------------")
		for _, eachNumber := range row.perRowCoordinates {
			//fmt.Println("looking at number")
			//fmt.Println(eachNumber)
			isValid := CheckCoordinate(eachNumber, gameData)
			//fmt.Println("checked coordinate")
			if isValid > 0 {
				fmt.Println("valid")
				fmt.Println(isValid)
				total += isValid
			}
			//fmt.Println("after row total is")
			//fmt.Println(total)
		}
		//fmt.Println("calculated all data")
	}
	//fmt.Println("total")
	//fmt.Println(total)
	return total
}

func CheckCoordinate(coordinates []int, gameData []RowElement) int {
	startCoordinate := []int{coordinates[1], coordinates[2]}
	endCoordinate := []int{coordinates[3], coordinates[4]}
	lengthOfLine := len(gameData[0].stringRow)
	//fmt.Println("coordinate 0")
	//fmt.Println(coordinates[0])

	//check element to the left if element is not 0
	if startCoordinate[0] > 0 {
		leftStartElement := startCoordinate[0] - 1
		if (checkIfContainsSymbol(gameData[startCoordinate[1]].stringRow[leftStartElement])) == true {
			//fmt.Println("left")
			return coordinates[0]
		}
	}

	//check element to the right of last element
	if startCoordinate[0] < lengthOfLine {
		rightEndElement := endCoordinate[0] + 1
		if (checkIfContainsSymbol(gameData[endCoordinate[1]].stringRow[rightEndElement])) == true {
			//fmt.Println("right")
			return coordinates[0]
		}
	}

	//check we're not on row 0
	if startCoordinate[1] != 0 {
		//check element above start element
		yCoordForRowAbove := startCoordinate[1] - 1
		if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[startCoordinate[0]])) == true {
			//fmt.Println("above first")
			return coordinates[0]
		}

		if (coordinates[4] - coordinates[2]) > 1 {
			//check element above middle element
			midCoordinate := startCoordinate[0] + 1
			if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[midCoordinate])) == true {
				//fmt.Println("above mid")
				return coordinates[0]
			}
		}

		//check element above last element
		if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[endCoordinate[0]])) == true {
			//fmt.Println("above last")
			return coordinates[0]
		}

		if startCoordinate[0] > 0 {
			//check element above and one to left of start element
			newXElement := startCoordinate[0] - 1
			if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[newXElement])) == true {
				//fmt.Println("above left of first")
				return coordinates[0]
			}
		}

		newXElement := startCoordinate[0] + 1
		//check element above and one to right of start element
		if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[newXElement])) == true {
			//fmt.Println("above right of first")
			return coordinates[0]
		}

		newXElement = endCoordinate[0] - 1
		//check element above and one to left of end element
		if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[endCoordinate[0]])) == true {
			//fmt.Println("above left of end")
			return coordinates[0]
		}

		if startCoordinate[0] < lengthOfLine {
			newXElement = endCoordinate[0] + 1
			//check element above and one to right of end element
			if (checkIfContainsSymbol(gameData[yCoordForRowAbove].stringRow[newXElement])) == true {
				//fmt.Println("above right of end")
				return coordinates[0]
			}
		}

	}

	//check we're not on final row
	if startCoordinate[1] < len(gameData)-1 {
		//check element below start element
		yCoordForRowBelow := startCoordinate[1] + 1
		if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[startCoordinate[0]])) == true {
			//fmt.Println("below start")
			return coordinates[0]
		}

		if (coordinates[4] - coordinates[2]) > 1 {
			//check element below middle element
			midCoordinate := startCoordinate[0] + 1
			if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[midCoordinate])) == true {
				//fmt.Println("below mid")
				return coordinates[0]
			}
		}

		//check element below last element
		if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[endCoordinate[1]])) == true {
			//fmt.Println("below last")
			return coordinates[0]
		}

		if startCoordinate[0] > 0 {
			//check element below and one to left of start element
			newXElement := startCoordinate[0] - 1
			if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[newXElement])) == true {
				//fmt.Println("below left of first")
				return coordinates[0]
			}
		}

		newXElement := startCoordinate[0] + 1
		//check element above and one to right of start element
		if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[newXElement])) == true {
			//fmt.Println("below right of first")
			return coordinates[0]
		}

		newXElement = endCoordinate[0] - 1
		//check element above and one to left of end element
		if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[endCoordinate[0]])) == true {
			//fmt.Println("below left of end")
			return coordinates[0]
		}

		if startCoordinate[0] < lengthOfLine {
			newXElement = endCoordinate[0] + 1
			//check element above and one to right of end element
			if (checkIfContainsSymbol(gameData[yCoordForRowBelow].stringRow[newXElement])) == true {
				//fmt.Println("below right of end")
				return coordinates[0]
			}
		}
	}

	return 0
}
