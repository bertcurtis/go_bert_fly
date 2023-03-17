package main

import (
	"fmt"

	"github.com/bertcurtis/go_bert_fly/internal"
)

func main() {
	me := internal.GetFlightList("deez")
	fmt.Printf("%s nutz\n", me)

}
