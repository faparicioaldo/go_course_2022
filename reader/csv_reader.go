package reader

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/faparicioaldo/bank-account-practice/entity"
)

type CSVReader struct{}

// It reads a CSV file in order to parsing and extract User data. The User structure is:
// Name: Persona name expected
// Email: valid email address
// Alias: alternate name to identify an user
func (csvReader *CSVReader) ReadFromFile(filePath string) (users []entity.AccountHolder, err error) {
	var reader *csv.Reader
	var content []byte
	content, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read from file %s", filePath)
	}
	if reader, err = csvReader.getCSVReader(content); err != nil {
		return nil, err
	}
	return csvReader.parseCSV(reader)
}

// This private function gets a pinter to csv.Reader depending on content param type
func (csvReader *CSVReader) getCSVReader(content interface{}) (users *csv.Reader, err error) {
	var reader io.Reader
	switch t := content.(type) {
	case []byte:
		reader = bytes.NewReader(content.([]byte))
	case string:
		reader = strings.NewReader(content.(string))
	default:
		return nil, fmt.Errorf("Could not define reader for given content %v", t)
	}

	return csv.NewReader(reader), nil
}

// Parses the CVS into a User structure
func (csvReader *CSVReader) parseCSV(reader *csv.Reader) (users []entity.AccountHolder, err error) {
	for {
		var record []string
		record, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("Error during file reading. Culd not read CSV. %v", err)
		}

		users = append(users, entity.AccountHolder{Name: record[0], Email: record[1], Alias: record[2]})
	}
	return users, nil
}
