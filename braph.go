package braph

import "log"

type Graph struct {
	cwidth int
	cells  []uint16
}

func NewGraph(width, height int) *Graph {
	this := Graph{
		cwidth: width / 2,
		cells:  make([]uint16, ((width+1)/2)*((height+3)/4)),
	}
	log.Printf("new(%d,%d) => [0..%d]", width, height, len(this.cells))
	return &this
}

func (this *Graph) Set(x, y int, on bool) {
	r := y / 4
	c := x / 2
	p := y%4 + 4*(x%2)
	//log.Printf("set(%d,%d,%v) => [%d+%d*%d] bit %d", x, y, on, r, this.cwidth, c, p)
	v := this.cells[c+r*this.cwidth]
	if on {
		v |= 1 << p
	} else {
		v &= 255 - 1<<p
	}
	this.cells[c+r*this.cwidth] = v
}

func (this *Graph) Strings() []string {
	out := []string{}
	buf := []rune{}
	for i, val := range this.cells {
		if i > 0 && i%this.cwidth == 0 {
			out = append(out, string(buf))
			buf = []rune{}
		}
		//log.Printf("%d: %v", i, val)
		buf = append(buf, chars[val])
	}
	return append(out, string(buf))
}
