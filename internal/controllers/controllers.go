package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bertcurtis/go_bert_fly/internal/models"
	"github.com/davecgh/go-spew/spew"
)

func SoulBrother(w http.ResponseWriter, r *http.Request) {
	flightInput := models.FlightsInput{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Unable to read request body")
		return
	}

	err = json.Unmarshal(body, &flightInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Request body is not valid. Valid input would be: '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'`)
		return
	}
	spew.Dump(body)

	spew.Dump(flightInput)
	// spew.Dump(w)

	w.WriteHeader(http.StatusOK)

}
