// +build ignore

package main

// gnuplot -p -e 'set grid; plot "gauss.txt" with steps'

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/go-hep/hist"
	"github.com/go-hep/hplot"
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

	// Make a plot and set its title.
	p, err := hplot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Gaussian"
	p.X.Label.Text = "X"
	//p.Y.Label.Text = "Y"

	hh, err := hplot.NewH1D(h)
	if err != nil {
		panic(err)
	}
	p.Add(hh)

	// Draw a grid behind the data
	p.Add(hplot.NewGrid())

	// Save the plot to a PDF file.
	if err := p.Save(6, 4, "hist.pdf"); err != nil {
		panic(err)
	}
}
