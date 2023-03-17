package main

import (
	"fmt"
	"net/http"

	"github.com/bertcurtis/go_bert_fly/internal/controllers"
)

func main() {
	http.Handle("/calculate", http.HandlerFunc(controllers.SoulBrother))
	fmt.Println("listening on localhost:8080/calculate")

	// ignoring the error value returned by ListenAndServe
	_ = http.ListenAndServe(":8080", nil)
}
