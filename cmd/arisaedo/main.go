package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/floydeconomy/arisaedo-go/cmd/arisaedo/node"
	"github.com/floydeconomy/arisaedo-go/cmd/arisaedo/utils"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &cli.App{
		Name: "Arisaedo",
		Version: "1.0.0",
		Usage: "Node of Arisaedo COVID-19 Data Aggregator",
		Copyright: "2020 Arisaedo <https://github.com/floydeconomy/>",
		Flags: []cli.Flag{
			&utils.NetworkFlag,
			&utils.ApiAddrFlag,
			&utils.ApiCorsFlag,
		},
		Action: Actions,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// todo: make this work
func Actions(ctx *cli.Context) error {
	// setup: exit
	exit := handleExitSignal()
	defer func() { log.Info("exited") }()

	// setup: api
	if err := utils.HandleAPIMainThread(ctx); err != nil {
		return err
	}

	// return
	return node.New().Run(exit)
}

func handleExitSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exitSignalCh := make(chan os.Signal)
		signal.Notify(exitSignalCh, os.Interrupt, os.Kill, syscall.SIGTERM)

		select {
		case sig:= <-exitSignalCh:
			log.Info("exit signal received", "signal", sig)
			cancel()
		}
	}()
	return ctx
}