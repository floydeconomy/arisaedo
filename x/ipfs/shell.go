package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
)

func IPFSConfigure() *shell.Shell {
	return shell.NewShell("https://ipfs.infura.io:5001")
}
