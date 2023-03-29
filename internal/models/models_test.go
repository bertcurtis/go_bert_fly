package models_test

import (
	"encoding/json"
	"testing"

	"github.com/bertcurtis/go_bert_fly/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFindInitialDeparture(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()
	assert := assert.New(t)

	sample := `[["ATL","JFK"],["PHX","ATL"],["SLC","PHX"], ["SFO","SLC"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(err)

	// arrL, depL, err := fi.ValidateLengths()
	// assert.Nil(err)
	// assert.Equal(4, arrL)
	// assert.Equal(4, depL)

	initial := fi.SolveInitial()
	assert.Equal([]string{"SFO", "SLC"}, initial)

	//assert.Equal(t, flightOutput.Result, "")
}

func TestFindNextDestination(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()
	assert := assert.New(t)

	sample := `[["ATL","JFK"],["PHX","ATL"],["SLC","PHX"], ["SFO","SLC"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(err)

	// arrL, depL, err := fi.ValidateLengths()
	// assert.Nil(err)
	// assert.Equal(4, arrL)
	// assert.Equal(4, depL)

	initial := fi.SolveNext()
	assert.Equal([]string{"SLC", "PHX"}, initial)

	//assert.Equal(t, flightOutput.Result, "")
}

func TestFindOrderedFlightPath(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()
	assert := assert.New(t)

	sample := `[["ATL","JFK"],["PHX","ATL"],["SLC","PHX"], ["SFO","SLC"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(err)

	// arrL, depL, err := fi.ValidateLengths()
	// assert.Nil(err)
	// assert.Equal(4, arrL)
	// assert.Equal(4, depL)

	orderedFlights, err := fi.SolveAll()
	assert.Nil(err)
	assert.Equal([]string{"SFO", "SLC"}, orderedFlights[0])
	assert.Equal([]string{"SLC", "PHX"}, orderedFlights[1])
	assert.Equal([]string{"PHX", "ATL"}, orderedFlights[2])
	assert.Equal([]string{"ATL", "JFK"}, orderedFlights[3])

	//assert.Equal(t, flightOutput.Result, "")
}

// func TestModelsSolve(t *testing.T) {
// 	// this tells go that this test can run in Parallel
// 	// with other t.parallel enabled unit tests
// 	t.Parallel()

// 	sample := `[["SFO","SLC"],["ATL","JFK"],["PHX","ATL"],["SLC","PHX"]]`
// 	fi := models.FlightsInput{}

// 	err := json.Unmarshal([]byte(sample), &fi)
// 	// ensure no Unmarshal error
// 	assert.Nil(t, err)

// 	flightOutput, err := fi.Solve()
// 	// should be no errors
// 	assert.Nil(t, err)

// 	assert.Equal(t, flightOutput.Result, "")
// }
