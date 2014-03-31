package main

type calogeo struct {
	ZMin  float64
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
	Calo.ZMin = 0.0
	Calo.ZMax = 1.0
	Calo.XYMin = -2.0
	Calo.XYMax = +2.0

	Calo.NbCellsInXY = 40
	Calo.NbLayers = 1

	// Cell size in x-y
	Calo.XYSize = (Calo.XYMax - Calo.XYMin) / float64(Calo.NbCellsInXY)
	Calo.ZSize = (Calo.ZMax - Calo.ZMin) / float64(Calo.NbLayers)
}

func (geo *calogeo) IsInside(x, y, z float64, addr *Cell) bool {
	*addr = Cell{}

	if x < geo.XYMin || x > geo.XYMax ||
		y < geo.XYMin || y > geo.XYMax ||
		z < geo.ZMin || z > geo.ZMax {
		return false
	}

	ix := int((x-geo.XYMin)/geo.XYSize) + 1
	iy := int((y-geo.XYMin)/geo.XYSize) + 1
	iz := int((z-geo.ZMin)/geo.ZSize) + 1
	*addr = Cell{X: ix, Y: iy, Layer: iz}
	return true
}

func (geo *calogeo) XCentre(addr *Cell) float64 {
	return geo.XYMin + (float64(addr.X-1)+0.5)*geo.XYSize
}

func (geo *calogeo) YCentre(addr *Cell) float64 {
	return geo.XYMin + (float64(addr.Y-1)+0.5)*geo.XYSize
}
