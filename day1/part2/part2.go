package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var col1, col2 []int

	// Open the file
	file, err := os.Open("../input1_a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		cols := strings.Fields(line)

		temp, _ := strconv.Atoi(cols[0])
		col1 = append(col1, temp)
		temp, _ = strconv.Atoi(cols[1])
		col2 = append(col2, temp)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(col1)
	sort.Ints(col2)

	var score int
	m := make(map[int]int)

	for _, val := range col2 {
		m[val]++
	}
	for _, c1 := range col1 {
		if val, ok := m[c1]; ok {
			score += c1 * val
		}
	}

	fmt.Printf("sim score total = %v\n", score)
}
