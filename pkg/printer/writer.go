package printer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/CbIPOKGIT/alkomag/pkg/details"
	customexcelwriter "github.com/CbIPOKGIT/custom-excel-writer"
)

type FileHeader struct {
	Name  string
	Title string
	Width int
}

var headers []FileHeader

func initWriter() *customexcelwriter.ExcelWriter {
	writer := new(customexcelwriter.ExcelWriter)
	writer.CreateFile()
	writer.CreateLonelySheet("Номенклатура")
	return writer
}

func printHeaders(writer *customexcelwriter.ExcelWriter) {
	for _, header := range getHeaders() {
		writer.SetCellValue(header.Title)
		writer.SetColumnsWidth(header.Width)
	}

	cells := &customexcelwriter.WorkBlock{
		RowStart: 1,
		RowEnd:   1,
		ColStart: 1,
		ColEnd:   writer.WorkBlock().ColEnd,
	}

	writer.ApplyStyle(writer.GetHeaderStyle(), cells)
}

func getHeaders() []FileHeader {
	if headers == nil {
		headers = []FileHeader{
			{Name: "index", Title: "#", Width: 5},
			{Name: "name", Title: "Номенклатура", Width: 40},
			{Name: "code", Title: "Артикул", Width: 25},
			{Name: "weight", Title: "Вага", Width: 20},
			{Name: "width", Title: "Ширина", Width: 20},
			{Name: "length", Title: "Довжина", Width: 20},
			{Name: "height", Title: "Висота", Width: 20},
		}
	}
	return headers
}

func printResult(result *details.RozetkaResult, writer *customexcelwriter.ExcelWriter, index int) {
	writer.CursorNextLine()

	name, articul := splitNameArticul(result.Name)

	for _, header := range getHeaders() {
		var value any

		switch header.Name {
		case "index":
			value = index
		case "name":
			value = name
		case "code":
			value = " " + articul
		case "weight":
			if result.Weight != 0 {
				value = fmt.Sprintf("%.2f", result.Weight)
			}
		case "width":
			if result.Width != 0 {
				value = fmt.Sprintf("%.2f", result.Width)
			}
		case "length":
			if result.Length != 0 {
				value = fmt.Sprintf("%.2f", result.Length)
			}
		case "height":
			if result.Height != 0 {
				value = fmt.Sprintf("%.2f", result.Height)
			}
		}

		writer.SetCellValue(value)
	}
}

func splitNameArticul(title string) (string, string) {
	parts := strings.Split(title, "(")
	if len(parts) == 1 {
		return strings.TrimSpace(parts[0]), ""
	}
	name, articul := strings.TrimSpace(parts[0]), parts[1]

	articul = regexp.MustCompile(`(?m)\D`).ReplaceAllString(articul, "")

	return name, articul
}

func applyCommonStyle(writer *customexcelwriter.ExcelWriter) {
	cells := &customexcelwriter.WorkBlock{
		RowStart: 2,
		RowEnd:   writer.WorkBlock().RowEnd,
		ColStart: 1,
		ColEnd:   writer.WorkBlock().ColEnd,
	}
	writer.ApplyStyle(writer.GetCommonStyle(), cells)
}
