package shell

func ShellConfigure() *shell.Shell {
	return shell.NewShell("https://ipfs.infura.io:5001")
}
