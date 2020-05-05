package store

import (
	"context"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
)

type Store struct {
	Node *core.IpfsNode
	Api coreiface.CoreAPI
}

func New(ctx context.Context) *Store {
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}

	cfg := &core.BuildCfg{
		Repo: r,
		Online: true,
	}

	nd, err := core.NewNode(ctx, cfg)
	if err != nil {
		panic(err)
	}

	api, err := coreapi.NewCoreAPI(nd)
	if err != nil {
		panic(err)
	}

	return &Store{
		Node: nd,
		Api: api,
	}
}