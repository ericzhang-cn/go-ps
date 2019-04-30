package main

import (
	"log"
	"os"

	"github.com/ericzhang-cn/go-ps/server"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	var (
		ip           string
		dataDir      string
		coordinators string
		logDir       string
	)

	app := cli.NewApp()
	app.Name = "server-cli"
	app.Version = "0.0.1"
	app.Usage = "starts server node"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "ip",
			Value:       "0.0.0.0:7771",
			Usage:       "ip address and port for binding",
			Destination: &ip,
		},
		cli.StringFlag{
			Name:        "data-dir",
			Value:       "/tmp/go-ps/data",
			Usage:       "directory of database",
			Destination: &dataDir,
		},
		cli.StringFlag{
			Name:        "coordinators",
			Value:       "127.0.0.1:7770",
			Usage:       "coordinators ip address and port",
			Destination: &coordinators,
		},
		cli.StringFlag{
			Name:        "log-dir",
			Value:       "/tmp/go-ps/logs",
			Usage:       "directory of log file",
			Destination: &logDir,
		},
	}

	app.Action = func(c *cli.Context) error {
		conf := server.Config{
			IPAddress: ip,
			BadgerDir: dataDir,
		}
		s := server.PsServer{
			Rank: 1,
			C:    &conf,
		}
		log.Printf("server listen to %s, node rank: %d", conf.IPAddress, s.Rank)
		if err := s.Serve(); err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
