package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	// Compute answers
	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

type Direction rune

const (
	Left    Direction = 'L'
	Right   Direction = 'R'
	Invalid Direction = 'I'
)

func updateLockPoint(line string, lockPoint int) (Direction, int, int, error) {
	rotations, err := strconv.Atoi(string(line[1:]))
	if err != nil {
		return Invalid, 0, 0, fmt.Errorf("bruh")
	}
	switch line[0] {
	case byte(Left):
		return Left, rotations, (lockPoint - rotations%100 + 100) % 100, nil
	case byte(Right):
		return Right, rotations, (lockPoint + rotations) % 100, nil
	default:
		return Invalid, 0, 0, fmt.Errorf("bruh")
	}
}

func solvePart1(lines []string) int {
	res := 0
	lockPoint := 50
	for _, line := range lines {
		_, _, newPoint, err := updateLockPoint(line, lockPoint)
		if err != nil {
			return 0
		}
		if newPoint == 0 {
			res += 1
		}
		lockPoint = newPoint
	}
	return res
}

func solvePart2(lines []string) int {
	lockPoint := 50
	res := 0
	for _, line := range lines {
		dir, rotation, newPoint, err := updateLockPoint(line, lockPoint)
		if err != nil {
			return 0
		}
		switch dir {
		case Right:
			res += (rotation + lockPoint) / 100
		case Left:
			if lockPoint == 0 {
				res += rotation / 100
			} else {
				res += (rotation + (100 - lockPoint)) / 100
			}
		}
		lockPoint = newPoint
	}
	return res
}
