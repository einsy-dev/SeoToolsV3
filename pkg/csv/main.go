package csv

import (
	"fmt"

	"github.com/einsy-dev/NetSail/internal/utils"
)

type CSV struct {
	Data [][]string
	cols map[string]int
	rows map[string]int
	key  int
}

func (c *CSV) Join(src *CSV) {
	join(c, src)
}

func (c *CSV) get(row string, col string) string {
	return c.Data[c.rows[row]][c.cols[col]]
}
func (c *CSV) set(row string, col string, val string) {
	c.Data[c.rows[row]][c.cols[col]] = val
}

func CSVFile(data string) (*CSV, error) {
	var csv = CSV{}
	var err error
	csv.Data, err = read(data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	csv.cols, csv.key = getCols(csv.Data)
	csv.rows = getRows(csv.Data, csv.key, func(k string) string { return utils.FormatDomain(k) })

	return &csv, nil
}
