package sort

import (
	"algorithms/result"
)

func SelectionSort(intSlice *[]int,order string) result.Result {
	switch order {
	case "ascending":
		for i := 0; i < len(*intSlice); i++ {
			minIndex := i
			for j := i + 1; j < len(*intSlice); j++ {
				if (*intSlice)[j] < (*intSlice)[minIndex] {
					minIndex = j
				}
			}
			(*intSlice)[i], (*intSlice)[minIndex] = (*intSlice)[minIndex], (*intSlice)[i]
		}

	case "descending":
		for i := 0; i < len(*intSlice); i++ {
			minIndex := i
			for j := i + 1; j < len(*intSlice); j++ {
				if (*intSlice)[j] > (*intSlice)[minIndex] {
					minIndex = j
				}
			}
			(*intSlice)[i], (*intSlice)[minIndex] = (*intSlice)[minIndex], (*intSlice)[i]
		}
	}
	return result.Result{BestCase: "O[N^2]", AverageCase: "O[N^2]", WorstCase: "O[N^2]", Space: "O[1]"}
}
