package csv

import "slices"

var keys = []string{"Item", "Domain", "Target"}

func getCols(data [][]string) (map[string]int, int) {
	var cols = make(map[string]int)
	var keyCol int

	for i, v := range data[0] {
		if slices.Contains(keys, v) {
			keyCol = i
			data[0][i] = "Domain"
			v = data[0][i]
		}
		cols[v] = i
	}
	return cols, keyCol
}
