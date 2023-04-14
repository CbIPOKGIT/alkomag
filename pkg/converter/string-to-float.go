package converter

import (
	"regexp"
	"strconv"
	"strings"
)

func StringToFloat(val string) float32 {
	var hasPoint bool
	var res string
	reg := regexp.MustCompile(`(?m)[\d\.,]`)

	for _, n := range strings.Split(val, "") {
		if !reg.Match([]byte(n)) {
			continue
		}

		if n == "." || n == "," {
			if hasPoint {
				break
			} else {
				hasPoint = true
				res += "."
			}
		} else {
			res += n
		}
	}

	if v, err := strconv.ParseFloat(res, 32); err == nil {
		return float32(v)
	} else {
		return 0
	}
}
