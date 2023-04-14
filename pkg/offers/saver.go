package offers

import (
	"encoding/json"
	"os"
)

func saveToFile(s []uint32) error {
	filename := "items.txt"

	j, err := json.Marshal(s)
	if err != nil {
		return nil
	}

	return os.WriteFile(filename, j, 0777)
}
