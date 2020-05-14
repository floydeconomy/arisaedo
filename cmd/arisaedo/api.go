package main

import (
	"context"
	"github.com/floydeconomy/arisaedo-go/api"
	"github.com/floydeconomy/arisaedo-go/co"
	"github.com/floydeconomy/arisaedo-go/store"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"github.com/urfave/cli/v2"
	"net"
	"net/http"
	"time"
)

func StartAPIServer(ctx *cli.Context, handler http.Handler) (string, func(), error) {
	addr := ctx.String(ApiAddrFlag.Name)

	// open connection
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return "", nil, errors.Wrapf(err, "listen API addr [%v]", addr)
	}

	// enable timeout
	timeout := ctx.Int(ApiTimeoutFlag.Name)
	if timeout > 0 {
		handler = handleAPITimeout(handler, time.Duration(timeout)*time.Millisecond)
	}

	// todo: middleware goes here

	// serve
	server := &http.Server{Handler: handler}
	var goes co.Goes
	goes.Go(func() {
		server.Serve(listener)
	})

	// return
	return "http://" + listener.Addr().String() + "/", func() {
		server.Close()
		goes.Wait()
	}, nil
}

func HandleAPIMainThread(ctx *cli.Context) error {
	handler, _ := api.New(ApiCorsFlag.Name)

	addr := ctx.String(ApiAddrFlag.Name)
	PrintAPIMessage(addr, ctx.String(NodeIDFlag.Name)) // todo: nodeID should check p2p comm

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

// todo: type check eth and ipfs strings
func HandleStore(ctx *cli.Context) *store.Store {
	ipfs := ctx.String(IPFSClientAddrFlag.Name)
	eth := ctx.String(EthClientAddrFlag.Name)
	s := store.New(store.Options{
		Db:    ipfs,
		Chain: eth,
	})
	PrintStoreMessage(ipfs, eth)
	return s
}

func handleAPITimeout(h http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func HandleAPIGoRoutine(ctx *cli.Context) error {
	handler, _ := api.New(ApiCorsFlag.Name)
	url, close, err := StartAPIServer(ctx, handler)
	if err != nil {
		return err
	}
	defer func() { log.Info("stopping API server...!"); close() }()
	PrintAPIMessage(url, "1") // todo: nodeID should check p2p comm
	return nil
}
