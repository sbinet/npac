package main

// Event holds informations about a calo event
type Event struct {
	Nbr    int     // event number
	McEne  float64 // true energy
	RecEne float64 // reconstructed energy

	Cells []CaloCell             // calorimeter cells
	Pos   struct{ X, Y float64 } // impact position
}
