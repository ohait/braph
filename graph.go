package braph

import (
	"bytes"
)

type LineGraph []float64

// display a graph with everything below the plot line filled
func (this LineGraph) RanderFillBelow(min, max float64) string {
	return this.Render(func(in float64) uint8 {
		b := int(4 * (in - min) / (max - min))
		//log.Printf("in: %v, b: %v", in, b)
		switch b {
		case 0:
			return 8
		case 1:
			return 4 + 8
		case 2:
			return 2 + 4 + 8
		case 3:
			return 1 + 2 + 4 + 8
		default:
			if b > 0 {
				return 1 + 2 + 4 + 8
			} else {
				return 0
			}
		}
	})
}

// render the graph by plotting a line between min and max
func (this LineGraph) RanderPlot(min, max float64) string {
	return this.Render(func(in float64) uint8 {
		b := int(4 * (in - min) / (max - min))
		//log.Printf("in: %v, b: %v", in, b)
		switch b {
		case 0:
			return 8
		case 1:
			return 4
		case 2:
			return 2
		case 3:
			return 1
		default:
			return 0
		}
	})
}

// render the graph using the given function that convert each entry to a bitmask
// bitmask should be between 0x00 and 0x0F
func (this LineGraph) Render(m func(float64) uint8) string {
	vals := this
	if len(vals) == 0 {
		return ""
	}

	// convert value to a bitmask

	var out bytes.Buffer
	for i := 0; i < len(vals); i += 2 {
		b := m(vals[i])
		if i+1 < len(vals) {
			b += m(vals[i+1]) * 16
		}
		out.WriteRune(chars[b])
	}
	return out.String()
}
