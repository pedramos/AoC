package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadByLine(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func cleanInput(input []string) [][4]int {
	// cleanInput := []int{}
	str := [2]string{}
	numsStr := [4]string{}
	nums := make([][4]int, len(input))
	for index, line := range input {
		str[0], str[1], _ = strings.Cut(line, ",")
		// fmt.Printf("%v\n", str)
		numsStr[0], numsStr[1], _ = strings.Cut(str[0], "-")
		numsStr[2], numsStr[3], _ = strings.Cut(str[1], "-")
		for i := 0; i < 4; i++ {
			nums[index][i], _ = strconv.Atoi(numsStr[i])
		}
	}
	return nums
}

func part1(input []string) int {
	nums := cleanInput(input)
	count := 0
	for i := range nums {
		if (nums[i][0] <= nums[i][2] && nums[i][1] >= nums[i][3]) || (nums[i][2] <= nums[i][0] && nums[i][3] >= nums[i][1]) {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	nums := cleanInput(input)
	count := 0
	for i := range nums {
		if (nums[i][2] <= nums[i][0] && nums[i][0] <= nums[i][3]) || (nums[i][2] <= nums[i][1] && nums[i][1] <= nums[i][3]) ||
			(nums[i][0] <= nums[i][2] && nums[i][2] <= nums[i][1]) || (nums[i][0] <= nums[i][3] && nums[i][3] <= nums[i][1]) {
			// fmt.Println(nums[i])
			count++
		}
	}
	return count
}

func main() {
	// input := ReadByLine("sample.txt")
	input := ReadByLine("input")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}