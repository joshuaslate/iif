package iif

import (
	"encoding/csv"
	"fmt"
	"io"
)

type IIFRow struct {
}

func parseHeaders(row []string) {

}

func parseIIF(rows [][]string) {
	for i := range rows {
		// The first cell of a row is always the "key"
		var key = rows[i][0]

		// If the row's first item starts with "!", it's a header row
		if rows[i][0][0:1] == "!" {
			key = rows[i][0][1:]
		}
	}
}

func FromFile(iifFile io.Reader) ([]IIFRow, error) {
	reader := csv.NewReader(iifFile)

	// IIFs are tab-delimited
	reader.Comma = '\t'

	// IIFs can have variable length
	reader.FieldsPerRecord = -1

	iifData, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read .iif file: %w", err)
	}

	parseIIF(iifData)

	return nil, nil
}
