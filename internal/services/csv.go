package services

import (
	"encoding/csv"
	"log"
	"slices"
	"strings"
)

type CSV struct{}

func (c *CSV) Parse(text []string) [][]string {
	data := csvData{}
	data.init()

	for _, v := range text {
		r := csv.NewReader(strings.NewReader(v))
		record, err := r.ReadAll()
		if err != nil {
			log.Fatalf("Error reading all records: %v", err)
		}
		d := csvData{}
		d.init()
		d.parse(record)
		data.join(&d)
	}
	return data.arr
}

type csvData struct {
	arr  [][]string
	cols map[string]int
	rows map[string]int
}

func (d *csvData) init() {
	d.arr = [][]string{}
	d.cols = make(map[string]int)
	d.rows = make(map[string]int)
}

var keys = []string{"Item", "Domain", "Target"}

func (d *csvData) parse(arr [][]string) {
	var kCol int = -1

	d.arr = arr

	for i, v := range arr[0] {
		if slices.Contains(keys, v) {
			kCol = i
			arr[0][kCol] = "Domain"
		}
		d.cols[v] = i
	}

	if kCol == -1 {
		log.Fatal("No key column fount")
		return
	}

	for i, v := range arr {
		d.rows[v[kCol]] = i
	}

}

func (d *csvData) join(data *csvData) {
	if len(d.arr) == 0 || len(d.arr[0]) == 0 {
		d.arr = data.arr
		d.rows = data.rows
		d.cols = data.cols
		return
	}
	d.cols = d.joinMeta(d.cols, data.cols)
	d.rows = d.joinMeta(d.rows, data.rows)

	maxRows := len(d.rows)
	maxCols := len(d.cols)

	d.arr = append(d.arr, make([][]string, maxRows-len(d.arr))...)
	for i := range d.arr {
		d.arr[i] = append(d.arr[i], make([]string, maxCols-len(d.arr[i]))...)
	}

	for r, i := range data.rows {
		for c, j := range data.cols {
			if data.arr[i][j] != "" {
				d.arr[d.rows[r]][d.cols[c]] = data.arr[i][j]
			}
		}
	}
}

func (d *csvData) joinMeta(dst map[string]int, src map[string]int) map[string]int {

	for k := range src {
		_, derr := dst[k]
		if !derr {
			dst[k] = len(dst)
			continue
		}
	}

	return dst
}
