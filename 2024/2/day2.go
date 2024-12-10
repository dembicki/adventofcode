package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("./inputs/input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	validSequences := 0
	validSequencesWithOneError := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, " ")
		numbers := []int{}
		for _, line := range lines {
			number, _ := strconv.Atoi(line)
			numbers = append(numbers, number)
		}

		if is_valid_sequence(numbers) {
			validSequences++
		}

		if is_valid_sequence_with_one_error(numbers) {
			validSequencesWithOneError++
		}
	}

	fmt.Println("Valid sequences:", validSequences)
	fmt.Println("Valid sequences with one error:", validSequencesWithOneError)
}

// Part 1
func is_valid_sequence(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff > 0 {
			decreasing = false
		} else {
			increasing = false
		}
	}

	return increasing || decreasing
}

// Part 2
func is_valid_sequence_with_one_error(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	if is_valid_sequence(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		temp := make([]int, 0)
		temp = append(temp, numbers[:i]...)
		temp = append(temp, numbers[i+1:]...)

		if is_valid_sequence(temp) {
			return true
		}
	}

	return false
}
