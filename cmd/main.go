package main

import (
	"fmt"
	"log"

	"github.com/Nidal-Bakir/hex/internal/adapters/core/arithmatic"
	_ "github.com/Nidal-Bakir/hex/internal/appenv" // autoload .env with init function. Do not remove this line
	"github.com/Nidal-Bakir/hex/internal/ports"
)

func main() {
	var arithmaticPort ports.ArithmaticPort
	arithmaticPort = arithmatic.NewAdapter()
	result, err := arithmaticPort.Add(6, 9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
