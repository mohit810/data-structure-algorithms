package sort

import (
	"algorithms/result"
)

func InsertSort(intSlice *[]int,order string) result.Result {
	switch order {
	case "ascending":
		for i := 1; i < len(*intSlice); i++ {
			key := (*intSlice)[i]
			j := i - 1

			for j > -1 && (*intSlice)[j] > key {
				(*intSlice)[j+1] = (*intSlice)[j]
				j -= 1
			}
			(*intSlice)[j+1] = key
		}
	case "descending":
		for i := 1; i < len(*intSlice); i++ {
			key := (*intSlice)[i]
			j := i - 1

			for j > -1 && (*intSlice)[j] < key {
				(*intSlice)[j+1] = (*intSlice)[j]
				j -= 1
			}
			(*intSlice)[j+1] = key
		}
	}

	return result.Result{BestCase: "O[N]", AverageCase: "O[N^2]", WorstCase: "O[N^2]", Space: "O[1]"}
}
