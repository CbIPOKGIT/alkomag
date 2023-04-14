package details

import (
	"regexp"
	"strings"

	"github.com/CbIPOKGIT/alkomag/pkg/converter"
)

type Sizes struct {
	values []float32
}

func NewSizeParser(rawSize string) *Sizes {
	var divider float32 = 1
	reg := regexp.MustCompile(`(?mi)мм`)
	if reg.Match([]byte(rawSize)) {
		divider = 10
	}

	reg = regexp.MustCompile(`(?m)[^\d\.,]`)
	parts := reg.Split(rawSize, -1)

	sizes := new(Sizes)
	sizes.values = make([]float32, 0, len(parts))

	for _, part := range parts {
		if floatValue := converter.StringToFloat(strings.TrimSpace(part)); floatValue > 0 {
			sizes.values = append(sizes.values, floatValue/divider)
		}
	}

	return sizes
}

func (s *Sizes) Len() int {
	return len(s.values)
}

func (s *Sizes) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s *Sizes) Less(i, j int) bool {
	return s.values[j] < s.values[i]
}

func (s *Sizes) GetWidth() float32 {
	return s.values[s.Len()-1]
}

func (s *Sizes) GetLength() float32 {
	if s.Len() <= 1 {
		return 0
	}
	if s.Len() > 2 || s.values[0] == s.values[1] {
		return s.values[s.Len()-2]
	} else {
		return 0
	}
}

func (s *Sizes) GetHeight() float32 {
	if s.Len() <= 1 {
		return 0
	}
	if s.Len() > 2 || s.values[0] != s.values[1] {
		return s.values[0]
	} else {
		return 0
	}
}
