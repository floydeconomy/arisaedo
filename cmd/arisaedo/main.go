package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/floydeconomy/arisaedo-go/api"
	"github.com/floydeconomy/arisaedo-go/cmd/arisaedo/node"
	"github.com/floydeconomy/arisaedo-go/cmd/arisaedo/utils"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"net"
	"net/http"
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
	if err := handleAPIEasy(ctx); err != nil {
		return err
	}

	// return
	return node.New().Run(exit)
}

func handleAPIEasy(ctx *cli.Context) error {
	handler, _ := api.New(utils.ApiCorsFlag.Name)

	addr := ctx.String(utils.ApiAddrFlag.Name)
	utils.PrintAPIMessage(addr, ctx.String(utils.NodeIDFlag.Name)) // todo: nodeID should check p2p comm

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrapf(err, "listen API addr [%v]", addr)
	}

	srv := &http.Server{Handler: handler}
	if err := srv.Serve(listener); err != nil {
		return errors.Wrapf(err, "serve API addr [%v]", addr)
	}
	return nil
}

func handleAPI(ctx *cli.Context) error {
	handler, _ := api.New(utils.ApiCorsFlag.Name)
	svrUrl, svrClose, err := utils.StartAPIServer(ctx, handler)
	if err != nil {
		return err
	}
	defer func() { log.Info("stopping API server...!"); svrClose() }()
	utils.PrintAPIMessage(svrUrl, "1") // todo: nodeID should check p2p comm
	return nil
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