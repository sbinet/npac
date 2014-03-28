package main

const (
	// calo dimensions, in meters.
	CaloZMin  = 0.0
	CaloZMax  = 1.0
	CaloXYMin = -2.0
	CaloXYMax = +2.0

	CaloNbCellsInXY = 40
	CaloNbLayers    = 1

	// Cell size in x-y
	CaloXYSize = (CaloXYMax - CaloXYMin) / float64(CaloNbCellsInXY)
	CaloZSize  = (CaloZMax - CaloZMin) / float64(CaloNbLayers)
)
