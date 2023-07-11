package braph_test

import (
	"math"
	"testing"

	"github.com/ohait/braph"
)

func TestMap(t *testing.T) {
	g := braph.NewMap(14, 11)
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

func TestMap2(t *testing.T) {
	g := braph.NewMap(100, 30)
	for x := 0; x < 100; x++ {
		g.Set(x, int(math.Sin((float64)(x)*0.2)*15+15), true)
		g.Set(x, int(math.Cos((float64)(x)*0.1)*15+15), true)
	}
	for _, s := range g.Strings() {
		t.Logf("»%s«", s)
	}
}

func TestGraph(t *testing.T) {
	g := braph.LineGraph{}
	for i := 1; i < 30; i++ {
		v := math.Sin(math.Pow(float64(i), 1.5) / 8)
		g = append(g, v)
		t.Log(g.RanderFillBelow(-.8, .8))
	}
}
