package utils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
)

func WriteJSON(data map[string]interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {

		return "", fmt.Errorf("error writing json: %s", err)
	}

	return string(b), nil
}

func WriteCSV(data map[string]interface{}) (string, error) {

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	header := []string{}
	row := []string{}

	for key, value := range data {
		header = append(header, key)
		row = append(row, fmt.Sprintf("%v", value))
	}

	err := writer.Write(header)
	if err != nil {
		return buf.String(), fmt.Errorf("error writing header csv: %s", err)
	}

	err = writer.Write(row)
	if err != nil {
		return buf.String(), fmt.Errorf("error writing rows csv: %s", err)
	}

	writer.Flush()

	return buf.String(), nil
}
