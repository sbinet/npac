package main

import (
	"fmt"
	"math"
	"math/rand"
)

type CaloSim struct {
	cells map[Cell]float64

	X0 float64 // radiation length (m)
	A  float64 // longitudinal shower parameter
	B  float64 // longitudinal shower parameter
	MR float64 // Moliere radius
}

// CaloData returns the calorimeter cells
func (sim *CaloSim) CaloData() []CaloCell {
	cells := make([]CaloCell, 0, len(sim.cells))
	for addr, ene := range sim.cells {
		cells = append(cells, CaloCell{Address: addr, E: ene})
	}
	return cells
}

//
func (sim *CaloSim) reset() {
	sim.cells = make(map[Cell]float64, Calo.NbCellsInXY*Calo.NbLayers)
}

// Simulate a shower of a given energy, starting from the impact point (x,y) of the electron
// at the front end of the calorimeter
func (sim *CaloSim) SimulateShower(x, y, ene float64) error {
	var err error

	sim.reset()

	// first check the point is within calorimeter volume
	if x < Calo.XYMin || x > Calo.XYMax ||
		y < Calo.XYMin || y > Calo.XYMax {
		return fmt.Errorf("invalid x,y position (%v,%v)", x, y)
	}

	xstart := x
	ystart := y
	zstart := Calo.ZMin + 0.001
	gmax := sim.Gamma(-1)

	const NbShootMax = 5000
	shootweight := 1.0 / NbShootMax
	for i := 0; i < NbShootMax; i++ {
		// generate in local coordinate system at starting point
		x1 := 0.0
		y1 := 0.0
		z1 := 0.0
		accept := false
		for !accept {
			z1 = rand.Float64() * (Calo.ZMax - Calo.ZMin)
			z := rand.Float64() * gmax
			if z < sim.Gamma(z1/sim.X0) {
				accept = true
			}
		}

		// generate transverse shape
		r := rand.NormFloat64() * sim.MR
		phi := rand.Float64() * 2 * math.Pi
		x1 = r * math.Cos(phi)
		y1 = r * math.Sin(phi)

		// translate to starting point
		x1 += xstart
		y1 += ystart
		z1 += zstart
		var addr Cell
		if Calo.IsInside(x1, y1, z1, &addr) {
			sim.cells[addr] += ene * shootweight
		}
	}

	return err
}

func (sim *CaloSim) Gamma(z float64) float64 {
	t := z / sim.X0
	// dE/dt ~ (m_b*t)**(a-1) * exp(-bt)
	dedt := math.Pow(sim.B*t, sim.A-1) * math.Exp(-sim.B*t)
	return dedt
}

func NewCaloSim() CaloSim {
	sim := CaloSim{
		cells: make(map[Cell]float64, Calo.NbCellsInXY*Calo.NbLayers),
		X0:    0.10, // 10cm
		A:     4,
		B:     0.5,
		MR:    0.05, // moliere radius: 5cm
	}
	return sim
}
