package main

import (
	"github.com/jon890/foscoin/cli"
	"github.com/jon890/foscoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
