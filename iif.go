package iif

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

// Entry contains the key ("TRNS", "SPL", etc.), raw row data, and sub entries as applicable (splits, in the case of transactions)
type Entry struct {
	Key     string
	RawData map[string]string
	Entries []Entry
}

func processRow(headers []string, rowData []string) map[string]string {
	mappedData := make(map[string]string)

	for i := range headers {
		mappedData[headers[i]] = rowData[i]
	}

	return mappedData
}

func parseIIF(rows [][]string) ([]Entry, error) {
	var entries []Entry
	headersByKey := make(map[string][]string)
	var activeTx *Entry

	for i := range rows {
		// The first cell of a row is always the "key"
		var key = rows[i][0]
		var cols = rows[i][1:]

		// If the row's first item starts with "!", it's a header row
		if rows[i][0][0:1] == "!" {
			key = rows[i][0][1:]

			if _, hasHeadersForKey := headersByKey[key]; !hasHeadersForKey {
				headersByKey[key] = cols
			}
		} else {
			if _, hasHeadersForKey := headersByKey[key]; !hasHeadersForKey {
				return nil, errors.New("invalid iif data, no headers found for key " + key)
			}

			switch key {
			case "TRNS":
				activeTx = &Entry{
					Key:     key,
					RawData: processRow(headersByKey[key], cols),
					Entries: []Entry{},
				}
				break
			case "SPL":
				if activeTx != nil {
					activeTx.Entries = append(activeTx.Entries, Entry{
						Key:     key,
						RawData: processRow(headersByKey[key], cols),
						Entries: nil,
					})
				}

				break
			case "ENDTRNS":
				if activeTx != nil {
					entries = append(entries, *activeTx)
					activeTx = nil
				}

				break
			default:
				entries = append(entries, Entry{
					Key:     key,
					RawData: processRow(headersByKey[key], cols),
					Entries: nil,
				})
			}
		}
	}

	return entries, nil
}

// FromFile accepts a tab-delimited Intuit Interchange Format (.IIF) file as input and parses it into a slice of Entry if valid
func FromFile(iiifile io.Reader) ([]Entry, error) {
	reader := csv.NewReader(iiifile)

	// IIFs are tab-delimited
	reader.Comma = '\t'

	// IIFs can have variable length
	reader.FieldsPerRecord = -1

	iifData, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read .iif file: %w", err)
	}

	return parseIIF(iifData)
}
