package main

import (
	"image/color"
	"math/rand"

	"github.com/go-hep/hist"
	//gnuplot "github.com/sbinet/go-gnuplot"
	"github.com/go-hep/hplot"
)

// run_sim simulates and fills the event
func run_sim(evt *Event) error {
	var err error

	mcene := 50.0
	evt.McEne = mcene // fixed true energy

	x := rand.Float64()*(Calo.XYMax-Calo.XYMin) + Calo.XYMin
	y := rand.Float64()*(Calo.XYMax-Calo.XYMin) + Calo.XYMin

	err = g_sim.SimulateShower(x, y, mcene)
	if err != nil {
		return err
	}

	evt.Cells = g_sim.CaloData()
	evt.Pos.X = x
	evt.Pos.Y = y

	return err
}

func ana_sim(evt *Event) error {
	var err error
	if evt.Nbr == 0 {
		h := hist.NewHist1D(100, 0., 100.)
		for _, cell := range evt.Cells {
			h.Fill(cell.E, 1.)
		}
		/*
			p, err := gnuplot.NewPlotter("", false, false)
			if err != nil {
				return err
			}
			xaxis := h.Axis()
			xdata := make([]float64, xaxis.Bins())
			ydata := make([]float64, xaxis.Bins())
			for i := range xdata {
				xdata[i] = xaxis.BinLowerEdge(i)
				ydata[i] = h.Content(i)
			}
			p.PlotXY(xdata, ydata, "Cells Energy")
			p.CheckedCmd("set terminal pdf; set output 'cell-ene.pdf';replot")
			p.CheckedCmd("q")
		*/
		p, err := hplot.New()
		if err != nil {
			return err
		}
		p.Title.Text = "cells energy"
		p.X.Label.Text = "E (GeV)"
		//p.Add(hplot.NewGrid())
		hh, err := hplot.NewH1D(h)
		if err != nil {
			return err
		}
		hh.Color = color.RGBA{R: 255, A: 255}
		p.Add(hh)
		err = p.Save(6, 4, "cells-ene.pdf")
		if err != nil {
			return err
		}
	}
	return err
}

// EOF
