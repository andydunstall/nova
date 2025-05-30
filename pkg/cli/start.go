package cli

func Start() error {
	cmd := newRootCommand()
	return cmd.Execute()
}
