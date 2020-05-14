package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &cli.App{
		Name:      "Arisaedo",
		Version:   "1.0.0",
		Copyright: "2020 Arisaedo <https://github.com/floydeconomy/>",
		Flags: []cli.Flag{
			&NetworkFlag,
			&ApiAddrFlag,
			&ApiCorsFlag,
			&IPFSClientAddrFlag,
			&EthClientAddrFlag,
		},
		Action: Actions,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Actions(ctx *cli.Context) error {
	// setup: exit
	exit := handleExitSignal()
	defer func() { log.Info("exited") }()

	// setup: api
	if err := HandleAPIGoRoutine(ctx); err != nil {
		return err
	}

	// setup: store
	s := HandleStore(ctx)

	// return
	return New(s).Run(exit)
}

func handleExitSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exitSignalCh := make(chan os.Signal)
		signal.Notify(exitSignalCh, os.Interrupt, os.Kill, syscall.SIGTERM)

		select {
		case sig := <-exitSignalCh:
			log.Info("exit signal received", "signal", sig)
			cancel()
		}
	}()
	return ctx
}
