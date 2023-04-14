package details

import (
	"encoding/json"
	"os"
)

func readIds() ([]uint32, error) {
	filename := "items.txt"

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	s := make([]uint32, 0)
	err = json.Unmarshal(content, &s)
	return s, err
}

func saveToFile(results []*RozetkaResult) error {
	data, err := json.Marshal(results)
	if err != nil {
		return err
	}

	var filename string = "results.txt"
	return os.WriteFile(filename, data, 0777)
}
