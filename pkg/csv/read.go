package csv

import (
	c "encoding/csv"
	"strings"
)

func read(v string) ([][]string, error) {
	r := c.NewReader(strings.NewReader(v))
	record, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return record, nil
}
