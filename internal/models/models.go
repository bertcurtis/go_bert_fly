package models

import "fmt"

type FlightsInput [][]string
type FlightOutput struct {
	Result string
}

func (fi FlightsInput) Solve() (fo FlightOutput, err error) {
	departures, arrivals, err := fi.splitFlightsInput()
	if err != nil {
		return fo, nil
	}
	findInitialDeparture(departures, arrivals)

	err = valdateInput(fi)
	return fo, err
}

func valdateInput(fi FlightsInput) error {
	return nil
}
func findInitialDeparture(departures, arrivals []string) string {
	initialDeparture := ""
	for _, departuresI := range departures {
		inArrivals := false
		for _, arrivalsI := range arrivals {
			if arrivalsI == departuresI {
				inArrivals = true
				break
			}
			if !inArrivals {
				initialDeparture = departuresI
			}

		}
	}
	return initialDeparture
}
func (fi FlightsInput) splitFlightsInput() (departures, arrivals []string, err error) {
	for _, flightPair := range fi {
		if len(flightPair) != 2 {
			return departures, arrivals, fmt.Errorf("Item %v does not have exactly 2 airports.", flightPair)
		}
		// this iterates through each item in the list and puts the first value in the first list,
		// and the second value in the second list turning a list of potentiallhy dozens of pairs into 2 lists
		// of departures and arrivals
		for i, airport := range flightPair {
			if i == 0 {
				departures = append(departures, airport)
			} else {
				arrivals = append(arrivals, airport)
			}

		}

	}
	return departures, arrivals, nil

}
