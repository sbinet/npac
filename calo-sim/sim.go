package main

import (
	"math/rand"
)

// run_sim simulates and fills the event
func run_sim(evt *Event) error {
	var err error

	mcene := 50.0
	evt.McEne = mcene // fixed true energy

	x := rand.Float64() * (Calo.XYMax-Calo.XYMin) + Calo.XYMin
	y := rand.Float64() * (Calo.XYMax-Calo.XYMin) + Calo.XYMin

	err = g_sim.SimulateShower(x,y,mcene)
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
	return err
}

// EOF
