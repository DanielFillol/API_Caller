package csv

import (
	"encoding/csv"
	"os"
)

// Read reads a CSV file from the given file path and returns a slice of strings
// containing the data from the first column of the file.
//	The separator parameter specifies the delimiter used in the CSV file.
//	The function returns an error if the file cannot be opened or read. It closes the file after reading it.
func Read(filePath string, separator rune) ([]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []string
	for _, line := range csvData {
		newLine := line[0]
		data = append(data, newLine)
	}

	return data, nil
}
