package main

type calogeo struct {
	ZMin float64
	ZMax  float64
	XYMin float64
	XYMax float64

	NbCellsInXY int
	NbLayers    int

	// Cell size in x-y
	XYSize float64
	ZSize  float64
}

var Calo calogeo

func init() {
	Calo.ZMin  = 0.0
	Calo.ZMax  = 1.0
	Calo.XYMin = -2.0
	Calo.XYMax = +2.0

	Calo.NbCellsInXY = 40
	Calo.NbLayers    = 1

	// Cell size in x-y
	Calo.XYSize = (Calo.XYMax - Calo.XYMin) / float64(Calo.NbCellsInXY)
	Calo.ZSize  = (Calo.ZMax - Calo.ZMin) / float64(Calo.NbLayers)
}

func (geo *calogeo) IsInside(x, y, z float64, addr *Cell) bool {
	*addr = Cell{}

	if x < Calo.XYMin || x > Calo.XYMax ||
		y < Calo.XYMin || y > Calo.XYMax ||
		z < Calo.ZMin || z > Calo.ZMax {
		return false
	}

	ix := int((x-Calo.XYMin)/Calo.XYSize)+1
	iy := int((y-Calo.XYMin)/Calo.XYSize)+1
	iz := int((z-Calo.ZMin)/Calo.ZSize)+1
	*addr = Cell{X:ix, Y:iy, Layer: iz}
	return true
}
