package calc

import (
	"math"
)

type Calculator struct {
	Mus string
	Aus string
}

func (c Calculator) CalPrecision(grouping, distance float64) float64 {
	mus := c.Mus
	aus := c.Aus

	if mus == "si" {
		if aus == "mil" {
			return math.Atan(grouping/(distance*100*2)) * 2 * 1000
		} else if aus == "moa" {
			return math.Atan(grouping/(distance*100*2)) * 2 * 60 * (180 / math.Pi)
		}
	} else if mus == "en" {
		if aus == "mil" {
			return math.Atan(grouping/(distance*36*2)) * 2 * 1000
		} else if aus == "moa" {
			return math.Atan(grouping/(distance*36*2)) * 2 * 60 * (180 / math.Pi)
		}
	}
	return 0
}

func (c Calculator) CalDistance(length, subtension float64) float64 {
	mus := c.Mus
	aus := c.Aus

	if mus == "si" {
		if aus == "mil" {
			return length / (math.Tan(subtension/(2*1000)) * 100 * 2)
		} else if aus == "moa" {
			return length / (math.Tan(subtension/(2*60*(180/math.Pi))) * 100 * 2)
		}
	} else if mus == "en" {
		if aus == "mil" {
			return length / (math.Tan(subtension/(2*1000)) * 36 * 2)
		} else if aus == "moa" {
			return length / (math.Tan(subtension/(2*60*(180/math.Pi))) * 36 * 2)
		}
	}
	return 0
}
