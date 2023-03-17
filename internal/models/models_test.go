package models_test

import (
	"encoding/json"
	"testing"

	"github.com/bertcurtis/go_bert_fly/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestModelsSolve(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["SFO","SLC"],["ATL","JFK"],["PHX","ATL"],["SLC","PHX"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput, err := fi.Solve()
	// should be no errors
	assert.Nil(t, err)

	assert.Equal(t, flightOutput.Result, "")
}
