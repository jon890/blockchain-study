package main

import (
	"github.com/jon890/foscoin/explorer"
	"github.com/jon890/foscoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
