# braph

simple drawing in golang using braille unicode 

## Map

Map are arbitrary size images where each pixel can be on or off (two color)

```go
  // create a new BitmaBitmap
	g := braph.NewMap(100, 30)

  // plot some dots
	for x := 0; x < 100; x++ {
    g.Set(x, 15, true) // plot the center line
		g.Set(
      x,
      int(math.Sin((float64)(x)*0.2)*15+15), // plot a sin centered at 15
      true)
	}

  // print it
	for _, s := range g.Strings() {
		fmt.Printf("⼁%s⼁\n", s)
	}
```

which outputs

```
⼁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡐⠉⠉⢂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠔⠉⠑⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠊⠉⠑⡀⠀⠀⠀⠀⼁
⼁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⠀⠀⠀⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠌⠀⠀⠀⠈⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠂⠀⠀⠀⠐⠀⠀⠀⠀⼁
⼁⠀⠀⠀⠀⠀⠀⠀⠀⢀⠁⠀⠀⠀⠀⠈⠄⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⠀⠀⠀⠀⠐⠀⠀⠀⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠀⠀⢁⠀⠀⠀⼁
⼁⣀⣀⣀⣀⣀⣀⣀⣀⣄⣀⣀⣀⣀⣀⣀⣐⣀⣀⣀⣀⣀⣀⣀⣀⣁⣀⣀⣀⣀⣀⣀⣁⣀⣀⣀⣀⣀⣀⣀⣐⣀⣀⣀⣀⣀⣀⣀⣄⣀⣀⼁
⼁⠐⠀⠀⠀⠀⠀⠀⠐⠀⠀⠀⠀⠀⠀⠀⠀⢁⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠀⠀⢀⠁⠀⠀⠀⠀⠀⠀⠀⠐⠀⠀⼁
⼁⠀⢁⠀⠀⠀⠀⠠⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠀⡐⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⡀⠀⠀⠀⠀⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠡⠀⼁
⼁⠀⠀⢂⠀⠀⢀⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠄⠀⠀⡠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⡀⠀⠀⠌⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢂⼁
⼁⠀⠀⠀⠑⠒⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠒⠒⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠒⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⼁
```

## LineGraph

Simple version of Map above, where a single text line is create from a series of values:

```go
  g := braph.LineGraph{}
  for i:=0; i<30; i++ {
		v := math.Sin(math.Pow(float64(i), 1.5) / 8)
    g = append(g, v)
  }
  fmt.Print(g.RenderFillBelow(-0.8, +0.8)) // use -0.8 and +0.8 as min/max
```
