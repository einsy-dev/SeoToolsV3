package csvParser

import (
	csv "github.com/einsy-dev/NetSail/pkg/csv"
)

type CsvParser struct{}

func (c *CsvParser) Parse(text []string) [][]string {
	var data csv.CSV

	for i, v := range text {
		c, err := csv.CSVFile(v)
		if err != nil {
			continue
		}
		if i == 0 {
			data = *c
		} else {
			data.Join(c)
		}
	}

	return data.Data

}
