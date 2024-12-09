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
	// Specify the path to the file you want to process
	filePath := "input1_a.txt"
	var col1, col2 []int

	// Open the file
	file, err := os.Open(filePath)
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

	var dist, total int
	for i, _ := range col1 {
		dist = col1[i] - col2[i]
		if dist < 0 {
			dist = dist * -1
		}
		fmt.Print(dist, "\n ")
		total += dist
	}

	fmt.Printf("distance total = %v\n", total)
}
