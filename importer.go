package topsearch

import (
	"encoding/csv"
	"strconv"
)

// Import returns a pair of index-mapped slices from the given io.Reader
func Import(r *csv.Reader) ([]string, []DataSet) {
	names := make([]string, 0)
	values := make([]DataSet, 0)

	for {
		record, err := r.Read()

		if err != nil {
			break
		}

		names = append(names, record[1])
		value, _ := strconv.ParseInt(record[0], 10, 64)
		d := DataSet { value }

		values = append(values, d)
	}

	return names, values
}
