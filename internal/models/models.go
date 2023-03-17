package models

type FlightsInput [][]string
type FlightOutput struct {
	Result string
}

func (fi FlightsInput) Solve() (fo FlightOutput, err error) {
	err = valdateInput(fi)
	return fo, err
}

func valdateInput(fi FlightsInput) error {
	return nil
}