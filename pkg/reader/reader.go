package reader

import (
	"encoding/csv"
	"os"
)

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename, seperator string) ([][]string, error) {
	basepath, err := os.Getwd()
	if err != nil {
		return [][]string{}, err
	}

	// Open CSV file
	f, err := os.Open(basepath + "/" + filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	csvreader := csv.NewReader(f)
	csvreader.Comma = []rune(seperator)[0]
	lines, err := csvreader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
