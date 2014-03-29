package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
)

var g_sim = NewCaloSim()

func main() {
	evtmax := 400
	flag.IntVar(&evtmax, "nevts", 400, "number of events to generate")

	flag.Parse()

	const fname = "event.gob"
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if enc == nil {
		panic(fmt.Errorf("could not create gob.Encoder"))
	}

	fmt.Printf(
		"::: npac-calo\n::: nevts=%d\n::: fname=%s\n",
		evtmax,
		fname,
	)

	for ievt := 0; ievt < evtmax; ievt++ {
		if ievt%100 == 0 {
			fmt.Printf(":: event=%d\n", ievt)
		}

		evt := Event{Nbr: ievt}

		// simulation
		err = run_sim(&evt)
		if err != nil {
			panic(err)
		}
		err = ana_sim(&evt)
		if err != nil {
			panic(err)
		}

		// reconstruction
		err = run_rec(&evt)
		if err != nil {
			panic(err)
		}

		err = enc.Encode(&evt)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("::: bye.\n")
}
