package sort

import (
	"algorithms/result"
)

func BubbleSort(intSlice *[]int, order string) result.Result {
	switch order {
	case "ascending":
		for i := 0; i < len(*intSlice)-1; i++ {
			for j := 0; j < len(*intSlice)-1-i; j++ {
				if (*intSlice)[j] > (*intSlice)[j+1] {
					(*intSlice)[j], (*intSlice)[j+1] = (*intSlice)[j+1], (*intSlice)[j]
				}
			}
		}

	case "descending":
		for i := 0; i < len(*intSlice)-1; i++ {
			for j := 0; j < len(*intSlice)-1-i; j++ {
				if (*intSlice)[j] < (*intSlice)[j+1] {
					(*intSlice)[j], (*intSlice)[j+1] = (*intSlice)[j+1], (*intSlice)[j]
				}
			}
		}

	}

	return result.Result{BestCase: "O[N]", AverageCase: "O[N^2]", WorstCase: "O[N^2]", Space: "O[1]"}
}
