package main

const (
	// calo dimensions, in meters.
	ZMin  = 0.0
	ZMax  = 1.0
	XYMin = -2.0
	XYMax = +2.0

	NbCellsInXY = 40
	NbLayers    = 1

	// Cell size in x-y
	XYSize = (XYMax - XYMin) / float64(NbCellsInXY)
	ZSize  = (ZMax - ZMin) / float64(NbLayers)
)
