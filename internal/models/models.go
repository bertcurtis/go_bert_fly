package models

import "fmt"

type FlightsInput [][]string
type FlightOutput struct {
	Result string
}

//	func (fi FlightsInput) Solve() (fo FlightOutput, err error) {
//		initial := findInitialDeparture(fi)
//		return initial
//	}
func (fi FlightsInput) SolveInitial() []string {
	initial := findDeparturePair(fi)

	return initial
}

func (fi FlightsInput) SolveNext() []string {
	initial := findDeparturePair(fi)
	next := findNextDestinationPair(fi, initial)
	return next
}

func findDeparturePair(fi FlightsInput) (returnPair []string) {
	fmt.Println("Length of flights input: ", len(fi))
	for _, flightPair := range fi {
		isDeparture := true
		fmt.Println("Input pair under iteration: ", flightPair)
		fmt.Println("Is this departure in: ", flightPair[0])
		for i := 0; i < len(fi); i++ {
			fmt.Println("arrval: ", fi[i][1])
			if flightPair[0] == fi[i][1] {
				isDeparture = false
				break
			}
		}
		if isDeparture {
			fmt.Println("found:", flightPair)
			returnPair = flightPair
			break
		}
	}
	return returnPair
}

func findNextDestinationPair(fi FlightsInput, flightPair []string) (returnPair []string) {
	for _, flightInputPair := range fi {
		if flightInputPair[0] == flightPair[1] {
			returnPair = flightInputPair
			break
		}
	}
	return returnPair
}

// func (fi FlightsInput) ValidateLengths() (arrivalsLength int, departuresLength int, err error) {
// 	departures, arrivals, err := fi.splitFlightsInput()
// 	if err != nil {
// 		return 0, 0, err
// 	}
// 	return len(arrivals), len(departures), err
// }

//	func valdateInput(fi FlightsInput) error {
//		return nil
//	}
// func findNextPair(fi FlightsInput, linkedDeparture string) {
// 	orderedFlights := [][]string{}
// 	for _, flightInputPair := range fi {
// 		fmt.Println("Input pair under iteration: ", flightInputPair)
// 		if flightInputPair[0] == linkedDeparture {
// 			orderedFlights = append(orderedFlights, flightInputPair)
// 			pairedArrival = flightInputPair[1]
// 		}
// 		orderedLen := len(orderedFlights)

// 		if orderedLen > 0 {
// 			linkedDeparture = orderedFlights[orderedLen-1][1]
// 		}
// 	}

// }
// func findInitialDeparture(fi FlightsInput) string {
// 	fmt.Println("Flights Input: ", fi)
// 	departures, arrivals, err := fi.splitFlightsInput()
// 	if err != nil {
// 		return "error"
// 	}
// 	initialDeparture := ""
// 	fmt.Println("departures: ", departures)
// 	fmt.Println("arrivals: ", arrivals)
// 	for _, departuresI := range departures {
// 		inArrivals := false
// 		fmt.Println("Comparing this departure: ", departuresI)
// 		for _, arrivalsI := range arrivals {
// 			if departuresI == arrivalsI {
// 				inArrivals = true
// 				fmt.Println("This departure is an arrival. This is not the initial departure", departuresI)
// 				break
// 			}
// 		}
// 		if !inArrivals {
// 			fmt.Println("There are no matching arrivals for this departure. Assuming this is the initial departure.", departuresI)
// 			initialDeparture = departuresI
// 			break
// 		}
// 	}
// 	return initialDeparture
// }

// func findArrivalByDeparture(fi FlightsInput, departure string) (arrival string) {
// 	for _, flightPair := range fi {
// 		if flightPair[0] == departure {
// 			arrival = flightPair[1]
// 			break
// 		}
// 	}
// 	return arrival
// }

// func addToOrderedList(fi FlightsInput, departure string) (arrival string){
// 	orderedFlights := [][]string{}
// 	linkedDeparture := initialDeparture
// 	for _, flightInputPair := range fi {
// 		fmt.Println("Input pair under iteration: ", flightInputPair)
// 		pairedArrival := ""
// 		if flightInputPair[0] == linkedDeparture {
// 			orderedFlights = append(orderedFlights, flightInputPair)
// 			pairedArrival = flightInputPair[1]
// 		}
// 		orderedLen := len(orderedFlights)

// 		if orderedLen > 0 {
// 			linkedDeparture = orderedFlights[orderedLen - 1][1]
// 		}
// 	}
// }

// func (fi FlightsInput) splitFlightsInput() (departures, arrivals []string, err error) {
// 	for _, flightPair := range fi {
// 		if len(flightPair) != 2 {
// 			return departures, arrivals, fmt.Errorf("Item %v does not have exactly 2 airports.", flightPair)
// 		}
// 		// this iterates through each item in the list and puts the first value in the first list,
// 		// and the second value in the second list turning a list of potentiallhy dozens of pairs into 2 lists
// 		// of departures and arrivals
// 		for i, airport := range flightPair {
// 			if i == 0 {
// 				departures = append(departures, airport)
// 			} else {
// 				arrivals = append(arrivals, airport)
// 			}

// 		}

// 	}
// 	return departures, arrivals, nil

// }
