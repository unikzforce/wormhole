package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vishvananda/netlink"
	"os"
	"strings"
	"wormhole/internal/switch_agent"
)

func main() {

	app := &cli.App{
		Name:  "switch_agent",
		Usage: "switch_agent, is the program that will reside in a network and will act as a switch device",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "interface-names",
				Value: "",
				Usage: "the name of the network interfaces to attach",
			},
		},
		Action: ActivateSwitchAgent,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("error %s", err)
		panic(err)
	}
}

func ActivateSwitchAgent(cCtx *cli.Context) error {
	networkInterfaces := findNetworkInterfaces(cCtx, "interface-names")
	switchAgent := switch_agent.NewSwitchAgent(networkInterfaces)

	return switchAgent.ActivateSwitchAgent()
}

func findNetworkInterfaces(cCtx *cli.Context, argumentName string) []netlink.Link {
	cliInterfaceNames := strings.TrimSpace(cCtx.String(argumentName))
	if cliInterfaceNames == "" {
		logrus.Fatal("--interface-names should be present and not empty")
	}

	logrus.Println(cliInterfaceNames)

	interfaceNames := strings.Split(cliInterfaceNames, ",")
	logrus.Println(interfaceNames)

	var ifaces []netlink.Link

	for _, ifaceName := range interfaceNames {
		iface, err := netlink.LinkByName(ifaceName)
		if err != nil {
			logrus.Fatalf("Getting interface %s: %s", ifaceName, err)
		}

		ifaces = append(ifaces, iface)
	}
	return ifaces
}
