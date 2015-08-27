package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/opsee/gozer/router"
)

const (
	version = "0.0.1"
)

func announce(c *cli.Context) {
	tableID := c.String("table")
	if tableID == "" {
		log.Fatal("Must specify --table")
	}

	rtr, err := router.NewRouter(tableID)
	if err != nil {
		log.Fatal(err.Error())
	}

	subnet := c.String("subnet")
	if subnet == "" {
		log.Fatal("Must specify --subnet")
	}

	err = rtr.Delete(subnet)
	if err != nil {
		log.Fatal(err)
	}

	err = rtr.Announce(subnet)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "router"
	app.Usage = "Announce a route for a subnet to this instance."
	app.Version = version
	app.Action = announce

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "subnet, s",
			Usage: "CIDR subnet to claim",
		},
		cli.StringFlag{
			Name:  "table, t",
			Usage: "Routing table ID",
		},
	}

	app.Run(os.Args)
}
