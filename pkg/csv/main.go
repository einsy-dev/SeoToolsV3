package csv

import "github.com/einsy-dev/NetSail/internal/utils"

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

func CSVFile(data string) CSV {
	var csv = CSV{}
	csv.Data = read(data)
	csv.cols, csv.key = getCols(csv.Data)
	csv.rows = getRows(csv.Data, csv.key, func(k string) string { return utils.FormatDomain(k) })
	return csv
}
