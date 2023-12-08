package Day_1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CreateData2() []string {
	var data []string
	file, _ := os.Open("day1pt2input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(strings.TrimSpace(scanner.Text()), "")
		convertedEntries := SubstituteWrittenNumbers(entries)
		data = append(data, convertedEntries)
	}

	_ = file.Close()
	return data
}

func GetAllIndexesOfElement(input string, value string) []int {
	returnableArray := []int{}
	re := regexp.MustCompile(value)
	elements := re.FindAllStringIndex(input, -1)
	for i := 0; i < len(elements); i++ {
		returnableArray = append(returnableArray, elements[i][0])
	}
	return returnableArray
}

func SubstituteWrittenNumbers(input []string) string {
	dataValues := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
	}

	justString := strings.Join(input, "")
	newString := map[int]string{}

	for k, v := range dataValues {
		if strings.Contains(justString, v) == true {
			indexes := GetAllIndexesOfElement(justString, v)
			for i := 0; i < len(indexes); i++ {
				newString[indexes[i]] = k
			}
		}
	}

	for i := 0; i < len(input); i++ {
		if _, err := strconv.Atoi(input[i]); err == nil {
			newString[i] = input[i]
		}
	}

	return GetFirstAndLastElement(newString)
}

func GetFirstAndLastElement(newString map[int]string) string {
	answer1, answer2 := "", ""
	lowestValue, highestValue := 10000, 0

	for k, v := range newString {
		if k <= lowestValue {
			lowestValue = k
			answer1 = v
		}
		if k >= highestValue {
			highestValue = k
			answer2 = v
		}
	}

	if answer1 == "" {
		answer1 = answer2
	}
	if answer2 == "" {
		answer2 = answer1
	}

	return answer1 + answer2
}

func SumUpAllElements2(elements []string) int {
	total := 0
	for i := 0; i < len(elements); i++ {
		convertedInt, _ := strconv.Atoi(elements[i])
		total += convertedInt
	}
	return total
}

func RunPt2() int {
	data := CreateData2()
	return SumUpAllElements2(data)
}
