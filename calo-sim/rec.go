package main

import (
	"math"
	"math/rand"
)

func run_rec(evt *Event) error {
	var err error
	mcene := evt.McEne                             // get true energy
	eReso := 0.1 * math.Sqrt(mcene)                // compute resolution
	eSmeared := mcene + eReso*(rand.Float64()-0.5) // smear with a flat distribution

	evt.RecEne = eSmeared // set the reconstructed energy
	return err
}
