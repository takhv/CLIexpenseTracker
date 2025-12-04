package storage

import (
	"encoding/json"
	"os"

	"github.com/takhv/expenseTracker/internal/model"
)

func readFile() ([]byte, error) {
	file := "data.json"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		os.Create(file)

		emptySlice := []model.Expense{}
		data, err := json.MarshalIndent(emptySlice, "", "\t")
		if err != nil {
			return data, err
		}

		os.WriteFile(file, data, 0644)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return data, err
	}
	return data, nil
}

func parseFromJSON(rawData []byte) ([]model.Expense, error) {
	var expenses []model.Expense
	err := json.Unmarshal(rawData, &expenses)
	if err != nil {
		return expenses, err
	}

	return expenses, nil
}

func ReadAndParse() ([]model.Expense, error) {
	rawData, err := readFile()
	if err != nil {
		return nil, err
	}
	data, err := parseFromJSON(rawData)
	if err != nil {
		return data, err
	}
	return data, nil
}

func ParseToJSONAndWrite(data []model.Expense) error {
	buildData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	file := "data.json"
	os.WriteFile(file, buildData, 0644)
	return nil
}
