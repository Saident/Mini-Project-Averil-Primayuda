package main

import (
	"github.com/Saident/Mini-Project-Averil-Primayuda/route"
)

func main() {
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
