package utils

import "github.com/urfave/cli/v2"

var (
	NetworkFlag = cli.StringFlag{
		Name: "network",
		Usage: "the network to join (main only supported)",
	}
	ApiCorsFlag = cli.StringFlag{
		Name: "api-cors",
		Value: "",
		Usage: "comma seperated list of domains from which to accept cross origin requests to API",
	}
	ApiAddrFlag = cli.StringFlag{
		Name: "api-addr",
		Value: "localhost:8669",
		Usage: "API service listening address",
	}
	ApiTimeoutFlag = cli.IntFlag{
		Name: "api-timeout",
		Value: 10000,
		Usage: "API request timeout value in milliseconds",
	}
	NodeIDFlag = cli.StringFlag{
		Name: "node-id",
		Value: "7214",
		Usage: "This should be removed eventually!",
	}
)