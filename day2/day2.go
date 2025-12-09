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
	// Specify the path to your text file
	filePath := "ranges.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	resultCounter := 0
	var totalOfIds uint64 = 0
	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		//length := len(line)

		ranges := strings.Split(line, ",")
		fmt.Printf("ranges: %v ", ranges)

		for _, v := range ranges {

			minAndMax := strings.Split(v, "-")
			base := 10 // Decimal base
			bitSize := 64
			min, err1 := strconv.ParseUint(minAndMax[0], base, bitSize)
			max, err2 := strconv.ParseUint(minAndMax[1], base, bitSize)

			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing string to unsigned int:", err)
				return
			}

			for i := min; i <= max; i++ {
				istring := strconv.FormatUint(uint64(i), 10)
				digits := len(istring)

				if digits%2 != 0 {
					continue
				}

				divider := math.Pow(10, float64(digits/2))

				firstHalf := i / uint64(divider)
				secondHalf := i % uint64(divider)
				// fmt.Printf("current number :%v divider: %v  firsthalf: %v Secondhalf: %v\n", i, divider, firstHalf, secondHalf)
				if firstHalf == secondHalf {
					//fmt.Printf("%v, ", i)
					totalOfIds += i
					resultCounter++
				}

			}

		}

	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file with scanner: %v\n", err)
	}

	fmt.Printf("We found this many: %v for a total of: %v", resultCounter, totalOfIds)
}
