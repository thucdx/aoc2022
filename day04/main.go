package main

import (
	"bufio"
	"fmt"
	"os"
)

func contain(s1, e1, s2, e2 int) bool {
	return (s1 >= s2 && e1 <= e2) || (s2 >= s1 && e2 <= e1)
}

func overlap(s1, e1, s2, e2 int) bool {
	//return (s1 <= s2 && s2 <= e1) || (s1 <= e2 && e2 <= e1) || (s2 <= s1 && e1 <= e2)
	return !(e1 < s2 || e2 < s1)
}

func main() {
	var path = "/Volumes/Data/workspace/learning/aoc2022/day04"
	fi, _ := os.Open(path + "/sample.inp")
	scanner := bufio.NewScanner(fi)

	var total = 0
	for scanner.Scan() {
		line := scanner.Text()
		var s1, e1, s2, e2 int
		//fmt.Println(line)
		fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
		if overlap(s1, e1, s2, e2) {
			total += 1
		}
		//if contain(s1, e1, s2, e2) {
		//	total += 1
		//}
	}

	fmt.Println(total)
}
