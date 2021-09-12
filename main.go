package main

import (
	"algorithms/sort"
	"algorithms/tree"
	"fmt"
)

func main() {

	slice := []int{2, 5, 3, 7, 4, 1, 9}

	fmt.Println("Slice -> ", slice)

	ascendingBubbleSort := sort.BubbleSort(&slice, "ascending")
	fmt.Println("Ascending order Bubble Sort -> ", slice, "\n Time Complexity -> \n Best -> ", ascendingBubbleSort.BestCase, "\n Average -> ", ascendingBubbleSort.AverageCase, "\n Worst -> ", ascendingBubbleSort.WorstCase, "\n Space Complexity -> ", ascendingBubbleSort.Space)

	descendingBubbleSort := sort.BubbleSort(&slice, "descending")
	fmt.Println("Descending order Bubble Sort -> ", slice, "\n Time Complexity -> \n Best -> ", descendingBubbleSort.BestCase, "\n Average -> ", descendingBubbleSort.AverageCase, "\n Worst -> ", descendingBubbleSort.WorstCase, "\n Space Complexity -> ", descendingBubbleSort.Space)

	ascendingInsertSort := sort.InsertSort(&slice, "ascending")
	fmt.Println("Ascending order Insert Sort -> ", slice, "\n Time Complexity -> \n Best -> ", ascendingInsertSort.BestCase, "\n Average -> ", ascendingInsertSort.AverageCase, "\n Worst -> ", ascendingInsertSort.WorstCase, "\n Space Complexity -> ", ascendingInsertSort.Space)

	descendingInsertSort := sort.InsertSort(&slice, "descending")
	fmt.Println("Descending order Insert Sort -> ", slice, "\n Time Complexity -> \n Best -> ", descendingInsertSort.BestCase, "\n Average -> ", descendingInsertSort.AverageCase, "\n Worst -> ", descendingInsertSort.WorstCase, "\n Space Complexity -> ", descendingInsertSort.Space)

	ascendingSelectionSort := sort.SelectionSort(&slice, "ascending")
	fmt.Println("Ascending order Selection Sort -> ", slice, "\n Time Complexity -> \n Best -> ", ascendingSelectionSort.BestCase, "\n Average -> ", ascendingSelectionSort.AverageCase, "\n Worst -> ", ascendingSelectionSort.WorstCase, "\n Space Complexity -> ", ascendingSelectionSort.Space)

	descendingSelectionSort := sort.SelectionSort(&slice, "descending")
	fmt.Println("Descending order Selection Sort -> ", slice, "\n Time Complexity -> \n Best -> ", descendingSelectionSort.BestCase, "\n Average -> ", descendingSelectionSort.AverageCase, "\n Worst -> ", descendingSelectionSort.WorstCase, "\n Space Complexity -> ", descendingSelectionSort.Space)

	result := sort.MergeSort(slice)
	fmt.Println("Ascending order Merge Sort -> ", result, "\n Time Complexity -> \n Best -> O[NlogN]", "\n Average -> O[NlogN]", "\n Worst -> O[NlogN]", "\n Space Complexity -> O[N]")

	ascendingQuickSort := sort.QuickSort(&slice)
	fmt.Println("Ascending order Quick Sort -> ", slice, "\n Time Complexity -> \n Best -> ", ascendingQuickSort.BestCase, "\n Average -> ", ascendingQuickSort.AverageCase, "\n Worst -> ", ascendingQuickSort.WorstCase, "\n Space Complexity -> ", ascendingQuickSort.Space)

	res, ascendingHeapSort := sort.HeapSort(slice)
	fmt.Println("Ascending order Heap Sort -> ", res, "\n Time Complexity -> \n Best -> ", ascendingHeapSort.BestCase, "\n Average -> ", ascendingHeapSort.AverageCase, "\n Worst -> ", ascendingHeapSort.WorstCase, "\n Space Complexity -> ", ascendingHeapSort.Space)

	fmt.Println("Binary Search Tree ->")
	inOrder, preOrder, postOrder, bstResult := tree.CreateBST(slice)
	fmt.Println(" InOrder -> ", inOrder, "\n PreOrder -> ", preOrder, "\n PostOrder -> ", postOrder)
	fmt.Println("\n Time Complexity -> \n Search --> \n Average -> ", bstResult.SearchAverageCase, "\n Worst -> ", bstResult.SearchWorstCase, "\n Insert --> \n Average -> ", bstResult.InsertAverageCase, "\n Worst -> ", bstResult.InsertWorstCase, "\n Delete --> \n Average -> ", bstResult.DeleteAverageCase, "\n Worst -> ", bstResult.DeleteWorstCase, "\n Space Complexity -> ", bstResult.Space)
}
