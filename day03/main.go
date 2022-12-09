package main

import (
	"bufio"
	"fmt"
	"os"
)

func process(text string) int {
	l := len(text)
	for i := 0; i < l/2; i++ {
		for j := l / 2; j < l; j++ {
			if text[i] == text[j] {
				if text[i] >= 'a' {
					return int(text[i]-'a') + 1
				} else {
					return int(text[i]-'A') + 27
				}

			}
		}
	}
	return 0
}

func main() {
	path := "/Volumes/Data/workspace/learning/aoc2022/day03"
	fi, _ := os.Open(path + "/input.inp")
	scanner := bufio.NewScanner(fi)

	total := 0
	group := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		group = append(group, text)
		if len(group) == 3 {
			val := processGroup(group)
			total += val
			group = make([]string, 0)
			fmt.Println(text, val)
		}
	}

	fmt.Println(total)
}

func processGroup(group []string) int {
	for i := 0; i < len(group[0]); i++ {
		for j := 0; j < len(group[1]); j++ {
			if group[0][i] == group[1][j] {
				for t := 0; t < len(group[2]); t++ {
					if group[2][t] == group[1][j] {
						c := group[1][j]
						if c >= 'a' {
							return int(c-'a') + 1
						} else {
							return int(c-'A') + 27
						}
					}
				}
			}
		}
	}
	return 0
}
