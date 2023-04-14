package offers

import (
	"fmt"
	"log"
)

const SEARCH_TEMPLATE = "https://xl-catalog-api.rozetka.com.ua/v4/goods/get?front-type=xl&country=UA&lang=ua&seller=rozetka&category_id=%s&page=%d"

func CollectOffers() error {

	mapa := make(map[uint32]any)

	categories := []string{
		"4649130", "4625409", "4649142", "4649136", "4649154", "4649166", "4649148", "4649178", "4649172", "4674322", "4649160",
		"4649052", "4649058", "4649064",
		"4626589", "4628313", "4649196",
	}

	for _, category := range categories {
		if err := collectCategoryOffers(category, mapa); err != nil {
			return err
		}
	}

	if err := saveToFile(mapaToSlice(mapa)); err != nil {
		return err
	}

	log.Printf("Saved %d elements", len(mapa))
	return nil
}

func collectCategoryOffers(category string, mapa map[uint32]any) error {
	total, page := 0, 0

	for {
		page++

		url := fmt.Sprintf(SEARCH_TEMPLATE, category, page)
		log.Println(url)
		resp, err := makeRequest(url)
		if err != nil {
			return err
		}

		for _, id := range resp.Data.Ids {
			mapa[id] = nil
		}

		total = resp.Data.TotalPages
		if page >= total {
			break
		}
	}
	return nil
}

func mapaToSlice(mapa map[uint32]any) []uint32 {
	s := make([]uint32, 0, len(mapa))

	for id := range mapa {
		s = append(s, id)
	}

	return s
}
