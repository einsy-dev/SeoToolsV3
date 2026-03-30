package csv

import (
	c "encoding/csv"
	"log"
	"strings"
)

func read(v string) [][]string {
	r := c.NewReader(strings.NewReader(v))
	record, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading all records: %v", err)
	}
	return record
}
