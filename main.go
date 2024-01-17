package main

import (
	"log"
	"os"

	"github.com/johnche/go-dns/dns"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-dns",
		Usage: "dns server in go",
		Action: func(*cli.Context) error {
			dns.Client()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("something went real wrong %v", err)
	}
}
