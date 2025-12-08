package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//The dial starts by pointing at 50.
	// read in each line and execute a rotation
	// R spins towards higher numbers
	//L spins towars lower numbers
	/*
		So, if the dial were pointing at 11, a rotation of R8 would cause the
		dial to point at 19. After that, a rotation of L19 would cause it to
		point at 0.

		Because the dial is a circle, turning the dial left from 0 one click makes
		it point at 99. Similarly, turning the dial right from 99 one click makes
		it point at 0.




		If the dial is at 40 and we rotate 233 to the left we will be at

		-233 % 99 = 64
		64 + 40 = 104
		104 % 99 = 5


		If the dial is at 40 and we rotate 233 to the right we will be at 75
		(233 % 99) = 35 -> 40 + 35 = 75
		233 % 99 = 35
		35 + 40 = 75
		75 % 99 = 5

		40-233 % 99 = 77
		abs(40-233) % 99 = 13

		100mod 99 = 1
		-100 mod 99 = 98

		if the dial is at 23 and we rotate 100 to the right we will be at 24
		if the dial is at 5 and we rotate 99 to the right we will be at 5


		if the dial is at 5 and we rotate 12 to the left we will be at 92
		(12 mod 99) = 12 -> 5 - 12 = -7
		if <0
			99 - 7 = 92



		count the number of times it hits zero

		-17mod4
		-17 - -20 = 3
	*/

	// Specify the path to your text file
	filePath := "rotation2.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	currPos := 50
	doorPassword := 0
	zeroCount := 0
	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		length := len(line)

		rotateBy, _ := strconv.Atoi(line[1:length])

		zeroCount += (currPos + rotateBy) / 100
		if string(line[0]) == "R" {
			currPos = ((rotateBy % 100) + currPos) % 100
		} else {
			// currPos = (negativeMod(rotateBy, 100) + currPos) % 100
			// maththing := ((-rotateBy) % 100)
			currPos = ((-rotateBy % 100) + currPos) % 100

			// fmt.Printf("NegativeMod: %v ", negativeMod(rotateBy, 100))
		}
		if currPos == 0 {
			doorPassword++
			zeroCount++
		}
		// fmt.Printf("currpos:%v\n", currPos)
	}
	fmt.Printf("doorpassword:%v\n", doorPassword)
	fmt.Printf("Zero Count :%v\n", zeroCount)
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file with scanner: %v\n", err)
	}
}

func negativeMod(a, b int) int {
	calc := -a % b
	if calc < 0 {
		return calc + b
	}
	return calc
}
