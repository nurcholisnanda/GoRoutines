package csv

import (
	"encoding/csv"
	"fmt"
	"goroutine/models"
	"os"
)

func ReadBeforeEodCsv(fileName string) ([]models.AfterEod, error) {

	//open csv file
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()

	//read file into records variable
	reader := csv.NewReader(f)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var afterEodData []models.AfterEod

	for index, record := range records {

		if index == 0 {
			continue
		}
		data, err := models.NewAfterEod(record)
		if err != nil {
			continue
		}
		afterEodData = append(afterEodData, *data)

	}

	return afterEodData, nil
}
