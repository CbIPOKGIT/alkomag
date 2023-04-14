package details

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/CbIPOKGIT/alkomag/pkg/converter"
)

const (
	PRODUCTNAME_URL    = "https://rozetka.com.ua/api/product-api/v4/goods/get-main?front-type=xl&country=UA&lang=ua&goodsId="
	CHARACTERISTIC_URL = "https://rozetka.com.ua/api/product-api/v4/goods/get-characteristic?front-type=xl&country=UA&lang=ua&goodsId="
)

type ProductnameResponse struct {
	Data struct {
		Title string `json:"title"`
	} `json:"data"`
}

type CharacteristicsResonse struct {
	Data []struct {
		Options []struct {
			Name   string `json:"name"`
			Values []struct {
				Title string `json:"title"`
			} `json:"values"`
		} `json:"options"`
	} `json:"data"`
}

func parseID(id uint32) (*RozetkaResult, error) {
	productname, err := getProductName(id)
	if err != nil {
		return nil, err
	}

	rawWeight, rawSize, err := getProductCharacteristics(id)
	if err != nil {
		return nil, err
	}

	result := &RozetkaResult{
		Id:        id,
		Name:      productname,
		RawSize:   rawSize,
		RawWeight: rawWeight,
	}

	if rawWeight != "" {
		result.Weight = converter.StringToFloat(rawWeight)
	}

	if rawSize != "" {
		result.Width, result.Height, result.Length = parseRawSize(rawSize)
	}

	return result, nil
}

func getProductName(id uint32) (string, error) {
	url := fmt.Sprintf("%s%d", PRODUCTNAME_URL, id)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := new(ProductnameResponse)
	if err := json.Unmarshal(content, data); err != nil {
		return "", err
	}

	return data.Data.Title, nil
}

func getProductCharacteristics(id uint32) (string, string, error) {
	url := fmt.Sprintf("%s%d", CHARACTERISTIC_URL, id)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	data := new(CharacteristicsResonse)
	if err := json.Unmarshal(content, &data); err != nil {
		return "", "", err
	}

	if len(data.Data) == 0 {
		return "", "", errors.New("empty_response_data")
	}

	var weight, size string

	for _, option := range data.Data[0].Options {
		if strings.Contains(option.Name, "razmery-") && len(option.Values) > 0 {
			size = strings.TrimSpace(option.Values[0].Title)
		}

		if strings.Contains(option.Name, "ves-") && len(option.Values) > 0 {
			weight = strings.TrimSpace(option.Values[0].Title)
		}
	}

	return weight, size, nil
}

func parseRawSize(rawSize string) (float32, float32, float32) {
	sizes := NewSizeParser(rawSize)
	sort.Sort(sizes)
	return sizes.GetWidth(), sizes.GetHeight(), sizes.GetLength()
}
