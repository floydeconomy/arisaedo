package utils

import (
	"context"
	"fmt"
	"github.com/floydeconomy/arisaedo-go/pkg/co"
	"github.com/pkg/errors"
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
	srv := &http.Server{Handler: handler}
	var goes co.Goes
	goes.Go(func() {
		srv.Serve(listener)
	})

	// return
	return "http://" +listener.Addr().String() + "/", func() {
		srv.Close()
		goes.Wait()
	}, nil
}


func handleAPITimeout(h http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func PrintAPIMessage(apiURL string, nodeID string) {
	fmt.Printf(`    API portal   [ %v ]
    Node ID      [ %v ]
`,
		apiURL,
		nodeID)
}