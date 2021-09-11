package sort

import "algorithms/result"

func QuickSort(intSlice *[]int) result.Result {

	sort(*intSlice, 0, len(*intSlice)-1)

	return result.Result{BestCase: "O[NlogN]", AverageCase: "O[NlogN]", WorstCase: "O[N^2]", Space: "O[logN]"}
}

func sort(intSlice []int, l,h int) {
	if l < h {
		j := partition(intSlice,l,h)
		sort(intSlice,l,j)
		sort(intSlice,j+1,h)
	}
}

func partition(arr []int, l,h int) int {
	pivot := arr[l]
	i, j:=l,h
	for i < j {
		if arr[i] <= pivot {
			i++
		}
		if arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j],arr[i]
		}
	}
	arr[l], arr[j] = arr[j],arr[l]
	return j
}