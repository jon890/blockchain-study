package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/jon890/foscoin/explorer"
	"github.com/jon890/foscoin/rest"
)

func usage() {
	fmt.Printf("Welcome to FoS 코인\n")
	fmt.Printf("Please use the following flags:\n")
	fmt.Printf("-port 4000:    Set the PORT of the server\n")
	fmt.Printf("-mode rest:    Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port)
		explorer.Start(*port + 1000)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
