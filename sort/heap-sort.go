package sort

import "algorithms/result"

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func parent(i int) int {
	return i / 2
}

func maxHeapify(a []int, i int) []int {

	//fmt.Printf("Array: %v\n", a)

	l := left(i) + 1
	r := right(i) + 1
	var largest int
	if l < len(a) && l >= 0 && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(a) && r >= 0 && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		//fmt.Printf("Exchanging: %d index (%d) with %d index (%d)\n", a[i], i, a[largest], largest)
		a[i], a[largest] = a[largest], a[i]
		a = maxHeapify(a, largest)
	}
	return a
}

func buildMaxHeap(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		//fmt.Printf("Building: %d index %d\n", a[i], i)
		a = maxHeapify(a, i)
	}
	return a
}

func HeapSort(a []int) ([]int,result.Result) {

	a = buildMaxHeap(a)
	//fmt.Printf("Starting sort ... array is %v\n", a)
	size := len(a)
	for i := size - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		size--
		maxHeapify(a[:size], 0)
	}
	return a, result.Result{BestCase: "O[NlogN]", AverageCase: "O[NlogN]", WorstCase: "O[NlogN]", Space: "O[1]"}
}