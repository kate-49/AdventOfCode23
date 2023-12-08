package Day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateData() [][]string {
	var data [][]string
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(strings.TrimSpace(scanner.Text()), "")
		data = append(data, entries)
	}

	_ = file.Close()

	return data
}

func GetFirstAndLastElementFromArray(intArray []int) []int {
	lengthOfArray := len(intArray)
	return []int{intArray[0], intArray[lengthOfArray-1]}
}

func SumUpAllElements(intArray [][]int) int {
	total := 0
	for i := 0; i < len(intArray); i++ {
		//condense these into strings to add together
		stringVariable := strconv.Itoa(intArray[i][0]) + strconv.Itoa(intArray[i][1])
		fmt.Println("string")
		fmt.Println(stringVariable)
		convertedInt, _ := strconv.Atoi(stringVariable)
		total += convertedInt
	}
	return total
}

func Run() int {
	data := CreateData()
	fmt.Println(data)
	numbers := [][]int{}

	for i := 0; i < len(data); i++ {
		element := []int{}
		for j := 0; j < len(data[i]); j++ {
			if convertedInt, err := strconv.Atoi(data[i][j]); err == nil {
				element = append(element, convertedInt)
			}
		}
		//if only one element in array append same element again
		if len(element) == 1 {
			element = append(element, element[0])
		}
		element = GetFirstAndLastElementFromArray(element)
		numbers = append(numbers, element)
	}

	return SumUpAllElements(numbers)
}
