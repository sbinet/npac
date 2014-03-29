package main

import (
	"fmt"
)

// Cell represents a calorimeter cell, using its position along the x, y and z axes
type Cell struct {
	X     int // cell index along the x axis
	Y     int // cell index along the y axis
	Layer int // cell index along the z axis
}

// valid address ?
func (cell Cell) IsValid() bool {
	if cell.X > Calo.NbCellsInXY ||
		cell.Y > Calo.NbCellsInXY ||
		cell.Layer > Calo.NbLayers ||
		cell.X < 0 || cell.Y < 0 || cell.Layer < 0 {
		return false
	}

	return true
}

// display cell address content
func (cell Cell) String() string {
	return fmt.Sprintf("(%d, %d, %d)", cell.X, cell.Y, cell.Layer)
}

// EOF
