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

	initial := fi.SolveInitial()
	assert.Equal([]string{"SFO", "SLC"}, initial)
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

	flightOutput := fi.SolveAll()
	assert.Equal(flightOutput.ErrorInformation, "")
	assert.Equal([]string{"SFO", "SLC"}, flightOutput.OrderedPath[0])
	assert.Equal([]string{"SLC", "PHX"}, flightOutput.OrderedPath[1])
	assert.Equal([]string{"PHX", "ATL"}, flightOutput.OrderedPath[2])
	assert.Equal([]string{"ATL", "JFK"}, flightOutput.OrderedPath[3])

}

func TestFaultyInputs(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()
	assert := assert.New(t)

	sample := `[["ATL","ATL"],["ATL","ATL"],["ATL","ATL"], ["ATL","ATL"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(err)

	initial := fi.SolveInitial()
	assert.Equal([]string{"SFO", "SLC"}, initial)
}
