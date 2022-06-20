package main

import (
	"fmt"

	"github.com/douglasmg7/imersaofsfc2-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "2",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
