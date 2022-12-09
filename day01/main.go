package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

var MAX_K int = 1

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) EnsureSize() {
	for h.Len() > MAX_K {
		heap.Pop(h)
	}
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	path := "/Volumes/Data/workspace/learning/aoc2022/day01"
	fi, _ := os.Open(path + "/input.inp")
	inp := bufio.NewReader(fi)
	defer fi.Close()
	// Set MAX_K = 1 for task1, MAX_K = 3 for task2
	MAX_K = 3
	h := &IntHeap{}
	heap.Init(h)

	scanner := bufio.NewScanner(inp)
	curStar := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println("Line", line)
		if line == "" {
			heap.Push(h, curStar)
			h.EnsureSize()
			curStar = 0

		} else {
			res, _ := strconv.Atoi(line)
			curStar += res
		}
	}

	heap.Push(h, curStar)
	h.EnsureSize()

	fmt.Println(sum(*h))
	//fmt.Println("------------")
	//for h.Len() > 0 {
	//	fmt.Println(heap.Pop(h))
	//}

	//fmt.Println((*h)[0:MAX_K], sum((*h)[0:MAX_K]))
}

func sum(arr []int) int {
	res := 0
	for _, val := range arr {
		res += val
	}
	return res
}
