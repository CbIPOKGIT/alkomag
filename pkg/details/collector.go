package details

import (
	"log"
	"time"
)

type RozetkaResult struct {
	Id        uint32  `json:"id"`
	Name      string  `json:"name"`
	Weight    float32 `json:"weight"`
	Length    float32 `json:"length"`
	Width     float32 `json:"width"`
	Height    float32 `json:"height"`
	RawSize   string  `json:"raw_size"`
	RawWeight string  `json:"raw_weight"`
}

func CollectData() error {
	ids, err := readIds()
	if err != nil {
		return err
	}

	results := make([]*RozetkaResult, len(ids))

	for i, id := range ids {
		log.Println(id)
		result, err := parseID(id)
		if err != nil {
			return err
		}

		results[i] = result

		if i%50 == 0 {
			log.Printf("%d results of %d - %d%%", i, len(ids), i*100/len(ids))
			time.Sleep(time.Second)
		}
	}

	return saveToFile(results)
}
