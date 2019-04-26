package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var port int
	var etcdEndpoints string
	var etcdNamespace string

	app := cli.NewApp()
	app.Name = "coordinator"
	app.Usage = "coordinate tasks for servers and workers"
	app.Description = "coordinator for go-ps"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port, p",
			Usage:       "binding to which port",
			Value:       7771,
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "etcd-endpoints, e",
			Usage:       "etcd endpoints ip address and port, format: 'ip1:port1,ip2:port2,...'",
			Value:       "127.0.0.1:2379",
			Destination: &etcdEndpoints,
		},
		cli.StringFlag{
			Name:        "etcd-namespace, n",
			Usage:       "namespace for etcd keys",
			Value:       "go-ps",
			Destination: &etcdNamespace,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
