package offers

import (
	"encoding/json"
	"io"
	"net/http"
)

type RozetkaResponse struct {
	Data struct {
		Ids        []uint32 `json:"ids"`
		TotalPages int      `json:"total_pages"`
	} `json:"data"`
}

func makeRequest(url string) (*RozetkaResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rr := new(RozetkaResponse)
	err = json.Unmarshal(data, rr)
	return rr, err
}
