package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "ns",
			Usage:   "根据host获取 name servers",
			Action:  func(c *cli.Context) error {
				ns, err := net.LookupNS(c.Args().Get(0))
				for _, val := range  ns{
					fmt.Println(val.Host)
				}
				return err
			},
		},
		{
			Name:    "cname",
			Usage:   "根据host获取 CNAME",
			Action:  func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.Args().Get(0))
				fmt.Println(cname)
				return err
			},
		},
		{
			Name:    "ip",
			Usage:   "根据host获取IP地址",
			Action:  func(c *cli.Context) error {
				ip, err := net.LookupIP(c.Args().Get(0))
				for _,val :=  range ip{
					fmt.Println(ip,val)
				}
				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}