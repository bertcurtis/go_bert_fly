package internal_test

import (
	"testing"

	"github.com/bertcurtis/go_bert_fly/internal"
	"github.com/stretchr/testify/assert"
)

func TestCalculateInvalidJSON(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()
	input := "deez"
	expected := "deez"
	strBody := internal.GetFlightList(input)
	assert.Equal(t, expected, strBody)
}
