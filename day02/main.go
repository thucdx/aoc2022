package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var m = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var v = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func scorePart2(first string, second string) int {
	score := 0
	switch second {
	case "X":
		{
			score += 0
			score += 1 + (v[first]+1)%3
		}
	case "Y":
		{
			score += 3
			score += v[first]
		}
	case "Z":
		{
			score += 6
			value := v[first] + 1
			if value > 3 {
				value = 1
			}
			score += value
		}

	}

	return score
}

func score(first string, second string) int {
	vF := v[first]
	vS := v[m[second]]

	score := vS
	if vF != vS {
		if vF < vS {
			if vS == vF+1 {
				// win
				score += 6
			} else {
				// lose
			}
		} else {
			if vF == vS+1 {
				// lose
			} else {
				// win
				score += 6
			}
		}
	} else {
		// draw
		score += 3
	}

	return score
}

func main() {
	path := "/Volumes/Data/workspace/learning/aoc2022/day02"
	fi, _ := os.Open(path + "/input.inp")
	//bfi := bufio.NewReader(fi)

	totalScore := 0
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		totalScore += scorePart2(parts[0], parts[1])
		fmt.Println(parts[0], parts[1], scorePart2(parts[0], parts[1]))
	}

	fmt.Println("Total Score", totalScore)
}
