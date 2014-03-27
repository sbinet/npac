package main

import (
	"fmt"
)

// CaloCell represents a calorimeter cell together with the energy deposit
type CaloCell struct {
	Address Cell // cell address
	E float64 // energy deposit
}

// display CaloCell
func (cell CaloCell) String() string {
	return fmt.Sprintf("[%s, %f]", cell.Address, cell.E)
}

// EOF
