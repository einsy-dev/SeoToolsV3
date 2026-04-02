package services

import (
	csv "github.com/einsy-dev/NetSail/pkg/csv"
)

type CSV struct{}

func (c *CSV) Parse(text []string) [][]string {
	var data csv.CSV

	for i, v := range text {
		var c = csv.CSVFile(v)
		if i == 0 {
			data = c
		} else {
			data.Join(&c)
		}

	}

	return data.Data

}
