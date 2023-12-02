package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

const DEFAULT_HOST = "apple.com"

// Generate will return cli app ready to exec
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Yole"
	app.Usage = "Skr skr gang gang"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: DEFAULT_HOST,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPS in internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search for server namein internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
