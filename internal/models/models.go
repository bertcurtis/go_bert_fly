package models

import "fmt"

type FlightsInput [][]string
type FlightOutput struct {
	InitialDeparturePair  []string
	FinalDepartureAirport string
	FinalArrivalAirport   string
	OrderedPath           [][]string
	ErrorInformation      string
}

// add handling for faulty inputs (duplicate entries, multiple destination and arrivals with mismatched pairs)
func (fi FlightsInput) SolveInitial() (fo FlightOutput) {
	initial, err := findInitialDeparturePair(fi)
	if err != nil {
		fo.ErrorInformation = err.Error()
		return fo
	}
	fo.InitialDeparturePair = initial
	return fo
}

func (fi FlightsInput) SolveAll() (fo FlightOutput) {
	initial, err := findInitialDeparturePair(fi)
	if err != nil {
		fo.ErrorInformation = err.Error()
		return fo
	}
	if len(fo.OrderedPath) < 1 {
		fo.OrderedPath = append(fo.OrderedPath, initial)
	}
	for i := 0; i < len(fi)-1; i++ {
		nextDestination, err := findNextDestinationPair(fi, fo.OrderedPath[len(fo.OrderedPath)-1])
		if err != nil {
			fo.ErrorInformation = err.Error()
			return fo
		}
		fo.OrderedPath = append(fo.OrderedPath, nextDestination)

	}
	if len(fo.OrderedPath) < len(fi) {
		err = fmt.Errorf("there are less flights in the ordered list than the inputs. This should be the same")
		fo.ErrorInformation = err.Error()

	}
	return fo
}

func findInitialDeparturePair(fi FlightsInput) (returnPair []string, err error) {
	// This will evaluate a set of flight paths and return a pair where the first item in the pair
	// does not have a match with any of the other items at the second ordinal position in the list of pairs
	for _, flightPair := range fi {
		isDeparture := true
		if len(flightPair) != 2 {
			return flightPair, fmt.Errorf("item %v does not have exactly 2 airports", flightPair)
		}

		for i := 0; i < len(fi); i++ {
			if flightPair[0] == fi[i][1] {
				isDeparture = false
				break
			}
		}
		if isDeparture {
			returnPair = flightPair
			break
		}
	}
	if len(returnPair) < 1 {
		return returnPair, fmt.Errorf("none of the flight pairs entered matched the criteria for an initial departure flight")
	}
	return returnPair, nil
}

func findNextDestinationPair(fi FlightsInput, flightPair []string) (returnPair []string, err error) {
	for _, flightInputPair := range fi {
		if flightInputPair[0] == flightPair[1] {
			returnPair = flightInputPair
		}
	}
	if len(returnPair) != 2 {
		return returnPair, fmt.Errorf("%v found when searching for next flight when there should be one pair", returnPair)
	}
	return returnPair, nil
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
