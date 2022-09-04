package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// an array is not heapify yet
	arrayForInit := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	arrayForPush := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}

	// apply Init mehtod
	maxHeap := buildHeapByInit(arrayForInit)

	// apply Push mehtod
	buildHeapByPush(arrayForPush)
	fmt.Println()

	// apply Pop method: remove and return the minimum element
	elementToRemove := heap.Pop(maxHeap)
	fmt.Println("removed element by Pop method: ", elementToRemove)
	fmt.Println()

	// apply Fix method
	fmt.Println("Current max heap: ", *maxHeap)
	// change number from 15 t0 50 at the index 2
	(*maxHeap)[2] = 50
	fmt.Println("Modify max heap value at index 2: ", *maxHeap)
	heap.Fix(maxHeap, 2)
	fmt.Println("Fix max heap: ", *maxHeap)
	fmt.Println()

	// apply Remove method
	fmt.Println("Current max heap: ", *maxHeap)
	// remove number at index 4
	heap.Remove(maxHeap, 4)
	fmt.Println("Current max heap after remove element in index 4: ", *maxHeap)
}

// Time: O(n)
func buildHeapByInit(array []int) *MaxHeap {
	// initialize the MinHeap that has implement the heap.Interface
	maxHeap := &MaxHeap{}
	*maxHeap = array
	heap.Init(maxHeap)
	fmt.Println("buildHeapByInit: ", *maxHeap)
	return maxHeap
}

// Time: O(nlogn)
func buildHeapByPush(array []int) *MaxHeap {
	// initialize the MinHeap that has implement the heap.Interface
	maxHeap := &MaxHeap{}
	for _, num := range array {
		heap.Push(maxHeap, num)
	}
	fmt.Println("buildHeapByPush: ", *maxHeap)
	return maxHeap
}

/* output
buildHeapByInit:  [91 31 40 25 10 15 1 9 22]
buildHeapByPush:  [91 40 31 25 10 15 1 9 22]

removed element by Pop method:  91

Current max heap:  [40 31 22 25 10 15 1 9]
Modify max heap value at index 2:  [40 31 50 25 10 15 1 9]
Fix max heap:  [50 31 40 25 10 15 1 9]

Current max heap:  [50 31 40 25 10 15 1 9]
Current max heap after remove element in index 4:  [50 31 40 25 9 15 1]
*/
