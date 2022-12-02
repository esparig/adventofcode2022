package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// An IntMaxHeap is a max-heap of ints.
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	var acum int = 0
	h := &IntMaxHeap{}
	heap.Init(h)

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		i, _ := strconv.Atoi(fileScanner.Text())
		acum += i
		if i == 0 {
			heap.Push(h, acum)
			acum = 0
		}
	}
	if acum > 0 {
		heap.Push(h, acum)
	}

	// Print top Elf calories carrier
	fmt.Printf("top: %d\n", (*h)[0])

	// Print top-3  Elves calories carrier sum
	top3 := heap.Pop(h).(int) + heap.Pop(h).(int) + heap.Pop(h).(int)
	fmt.Printf("top3: %d\n", top3)

	readFile.Close()

}
