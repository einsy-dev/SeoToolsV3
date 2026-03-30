package csv

func getRows(data [][]string, key int, format func(string) string) map[string]int {
	var rows = make(map[string]int)

	for i, v := range data {
		nKey := v[key]
		if format != nil {
			nKey = format(v[key])
			v[key] = nKey
		}
		rows[nKey] = i
	}

	return rows
}
