package csv

import (
	"encoding/csv"
	"fmt"
	"goroutine/models"
	"os"
	"strconv"
)

func WriteAfterEodCsv(fileName string, afterEodData []models.AfterEod) error {

	//open csv file
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModeType)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Comma = ';'

	defer writer.Flush()

	titleRow := []string{
		"id",
		"Nama",
		"Age",
		"Balanced",
		"No 2b Thread-No",
		"No 3 Thread-No",
		"Previous Balanced",
		"Average Balanced",
		"No 1 Thread-No",
		"Free Transfer",
		"No 2a Thread-No",
	}
	writer.Write(titleRow)

	for _, afterEodDatum := range afterEodData {
		record := []string{
			strconv.Itoa(afterEodDatum.ID),
			afterEodDatum.Nama,
			strconv.Itoa(afterEodDatum.Age),
			strconv.Itoa(afterEodDatum.Balanced),
			afterEodDatum.No2BThreadNo,
			afterEodDatum.No3ThreadNo,
			strconv.Itoa(afterEodDatum.PreviousBalanced),
			strconv.FormatFloat(float64(afterEodDatum.AverageBalanced), 'g', -1, 32),
			afterEodDatum.No1ThreadNo,
			strconv.Itoa(afterEodDatum.FreeTransfer),
			afterEodDatum.No2AThreadNo,
		}
		writer.Write(record)
	}

	return nil
}
