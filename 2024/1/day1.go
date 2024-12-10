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
	file, _ := os.Open("./inputs/input.txt")

	list1 := []int{}
	list2 := []int{}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		num1, _ := strconv.Atoi(line[0])
		num2, _ := strconv.Atoi(line[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// Part 1
	sum1 := 0
	for i := 0; i < len(list1); i++ {
		sum1 += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(sum1)

	// Part 2
	sum2 := 0
	for i := 0; i < len(list1); i++ {
		count := 0
		for _, num := range list2 {
			if list1[i] == num {
				count++
			}
		}
		sum2 += list1[i] * count
	}

	fmt.Println(sum2)
}
