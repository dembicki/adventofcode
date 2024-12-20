package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.Open("./inputs/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	sum := 0
	regex := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(regex)

	// Read all lines from the file
	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindAllStringSubmatch(text, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	fmt.Println(sum)
}
