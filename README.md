# braph

simple drawing in golang using braille unicode 

```go
  // create a new Graph
	g := braph.NewGraph(100, 30)

  // plot some dots
	for x := 0; x < 100; x++ {
		g.Set(x, int(math.Sin((float64)(x)*0.2)*15+15), true)
	}

  // print it
	for _, s := range g.Strings() {
		fmt.Printf("⼁%s⼁\n", s)
	}
```

