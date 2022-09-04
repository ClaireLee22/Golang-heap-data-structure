package main

import "fmt"

type MinHeap []int

// Time: O(n) | Space: O(1)
// heapify an array
func (h *MinHeap) BuildHeap(array []int) {
	lastNonLeafNodeIdx := (len(array) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(array) - 1
		h.siftDown(currentIdx, endIdx) // siftDown is more efficient than siftUp
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its smaller child node until it is in the correct position
func (h *MinHeap) siftDown(currentIdx int, endIdx int) {
	leftChildIdx := currentIdx*2 + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx*2 + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		// get the smaller child node to swap
		idxToSwap := leftChildIdx
		if rightChildIdx != -1 && (*h)[rightChildIdx] < (*h)[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check if value of swap node is less than the value at currentIdx
		if (*h)[idxToSwap] < (*h)[currentIdx] {
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
func (h *MinHeap) siftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && (*h)[currentIdx] < (*h)[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Time: O(logn) | Space: O(1)
// insert a new value to the end of the tree and update heap ordering
func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

// Time: O(logn) | Space: O(1)
// remove and return the minimum value and update heap ordering
func (h *MinHeap) Remove() int {
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
func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := &MinHeap{}
	*minHeap = array
	minHeap.BuildHeap(array)
	fmt.Println("build min heap: ", *minHeap)

	// apply Remove method
	valueToRemove := minHeap.Remove()
	fmt.Println("root value: ", valueToRemove)
	fmt.Println("min heap after Remove root: ", *minHeap)

	// apply Insert method
	// append a new value, 2
	minHeap.Insert(2)
	fmt.Println("min heap after insert value 2: ", *minHeap)
}

/* output
build min heap:  [1 10 9 22 31 15 40 25 91]
root value:  1
min heap after Remove root:  [9 10 15 22 31 91 40 25]
min heap after insert value 2:  [2 9 15 10 31 91 40 25 22]
*/
