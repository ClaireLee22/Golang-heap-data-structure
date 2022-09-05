package main

import "fmt"

type MaxHeap []int

// Time: O(n) | Space: O(1)
// heapify an array
func (h *MaxHeap) BuildHeap(array []int) {
	lastNonLeafNodeIdx := (len(array) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(array) - 1
		h.siftDown(currentIdx, endIdx) // siftDown is more efficient than siftUp
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its smaller child node until it is in the correct position
func (h *MaxHeap) siftDown(currentIdx int, endIdx int) {
	leftChildIdx := currentIdx*2 + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx*2 + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		// get the larger child node to swap
		idxToSwap := leftChildIdx
		if rightChildIdx != -1 && (*h)[rightChildIdx] > (*h)[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check if value of swap node is greater than the value at currentIdx
		if (*h)[idxToSwap] > (*h)[currentIdx] {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftChildIdx = currentIdx*2 + 1

		} else {
			return
		}
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its parent node until it is in the correct position
func (h *MaxHeap) siftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && (*h)[currentIdx] > (*h)[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Time: O(logn) | Space: O(1)
// insert a new value to the end of the tree and update heap ordering
func (h *MaxHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

// Time: O(logn) | Space: O(1)
// remove and return the minimum value and update heap ordering
func (h *MaxHeap) Remove() int {
	n := len(*h)
	// swap the first element and the last element in the array
	h.swap(0, n-1)
	valueToRemove := (*h)[n-1]
	// pop the last element in the array
	*h = (*h)[:n-1]
	// call siftDown to update heap ordering
	h.siftDown(0, n-2)

	return valueToRemove
}

// Time: O(1) | Space: O(1)
func (h MaxHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := &MaxHeap{}
	*minHeap = array
	minHeap.BuildHeap(array)
	fmt.Println("build max heap: ", *minHeap)

	// apply Remove method
	valueToRemove := minHeap.Remove()
	fmt.Println("root value: ", valueToRemove)
	fmt.Println("max heap after Remove root: ", *minHeap)

	// apply Insert method
	// append a new value, 50
	minHeap.Insert(50)
	fmt.Println("max heap after insert value 50: ", *minHeap)
}

/* output
build max heap:  [91 31 40 25 10 15 1 9 22]
root value:  91
max heap after Remove root:  [40 31 22 25 10 15 1 9]
max heap after insert value 50:  [50 40 22 31 10 15 1 9 25]
*/
