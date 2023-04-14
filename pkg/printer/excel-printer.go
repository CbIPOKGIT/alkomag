package printer

import (
	"encoding/json"
	"os"

	"github.com/CbIPOKGIT/alkomag/pkg/details"
	customexcelwriter "github.com/CbIPOKGIT/custom-excel-writer"
)

func PrintResults() error {
	results, err := loadData()
	if err != nil {
		return err
	}

	writer := initWriter()

	printHeaders(writer)

	for index, result := range results {
		printResult(result, writer, index+1)
	}

	applyCommonStyle(writer)

	return saveToFile(writer)
}

func loadData() ([]*details.RozetkaResult, error) {
	content, err := os.ReadFile("results.txt")
	if err != nil {
		return nil, err
	}

	results := make([]*details.RozetkaResult, 0)
	return results, json.Unmarshal(content, &results)
}

func saveToFile(writer *customexcelwriter.ExcelWriter) error {
	writer.AlignFileRows()
	return writer.SaveFile("results.xlsx")
}
