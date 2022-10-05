package utils

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"log"
)

// Parse Map
type Record struct {
	Key       string            `csv:"Key"`
	OtherData map[string]string `csv:"-"`
}

func ParseMap(csvFile io.Reader) (map[string]map[string]string, error) {
	r := csv.NewReader(csvFile)
	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		return nil, err
	}

	m := make(map[string]map[string]string)
	header := dec.Header()
	for {
		row := Record{OtherData: make(map[string]string)}
		if err := dec.Decode(&row); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for _, i := range dec.Unused() {
			value := dec.Record()[i]
			row.OtherData[header[i]] = value

		}

		m[row.Key] = row.OtherData
	}

	return m, nil
}
