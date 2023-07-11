package braph_test

import (
	"math"
	"testing"

	"github.com/ohait/braph"
)

func TestBraph(t *testing.T) {
	g := braph.NewGraph(14, 11)
	for i := 0; i < 11; i++ {
		g.Set(0, i, true)
		g.Set(12, i, true)
		g.Set(1+i, i, true)
		g.Set(11-i, i, true)
	}
	for _, s := range g.Strings() {
		t.Logf("»%s«", s)
	}
	// TODO properly test it, for now it just should look pretty
}

func TestBraph2(t *testing.T) {
	g := braph.NewGraph(100, 30)
	for x := 0; x < 100; x++ {
		g.Set(x, int(math.Sin((float64)(x)*0.2)*15+15), true)
		g.Set(x, int(math.Cos((float64)(x)*0.1)*15+15), true)
	}
	for _, s := range g.Strings() {
		t.Logf("»%s«", s)
	}
}
