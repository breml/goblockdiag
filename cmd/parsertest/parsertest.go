package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/breml/blockdiag"
)

func main() {
	simple := ""
	simple =
		`blockdiag {
	A -> B -> C -> D -> E;
	A -> F -> C -> G -> E;
}`

	got, err := blockdiag.ParseReader("simple.diag", strings.NewReader(simple))
	if err != nil {
		log.Fatal("Parse error:", err)
	}
	diag := got.(blockdiag.Diag)

	diag.PlaceInGrid()
	fmt.Printf("%s\n", diag.String())
}
