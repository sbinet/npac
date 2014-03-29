// +build ignore

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/go-hep/hist"
)

func main() {
	h := hist.NewH1DST("gaus", 100, -2., 2.)
	f, err := os.Create("gauss.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mu := 0.0
	sigma := 1.0
	frac := 1 / (sigma * math.Sqrt(2*math.Pi))
	// Draw 10000 numbers according the gaussian law
	// We use rand.Float64()*MAX  to draw numbers
	// uniformly  between  0. and MAX
	for i := 0; i < 10000; {
		x := rand.Float64()*4 - 2
		y := frac * math.Exp(-((x-mu)*(x-mu))/(2*sigma*sigma))
		ymc := rand.Float64()
		if ymc <= y {
			// we  fill here the histogram
			h.Fill(x, 1.0)
			i++
			//fmt.Fprintf(f, "%e\n", x)
		}
	}
	fmt.Printf("entries= %d\n", h.Entries())
	fmt.Printf("mean=    %8.3f\n", h.Mean())
	fmt.Printf("RMS=     %8.3f\n", h.RMS())

	ax := h.Axis()
	for i := 0; i < int(ax.Bins()); i++ {
		x := ax.BinLowerEdge(i)
		y := h.Content(i)
		fmt.Fprintf(f, "%e %e\n", x, y)
	}
}
