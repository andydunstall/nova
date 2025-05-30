package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build path [flags]",
		Short: "build a Nova program",
		Long:  `...`,
	}

	cmd.Run = func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			exitError(fmt.Errorf("build: missing path"))
		}
		if len(args) > 1 {
			exitError(fmt.Errorf("build: only one path is supported"))
		}

		if err := runBuild(args[0]); err != nil {
			exitError(fmt.Errorf("build: %w", err))
		}
	}

	return cmd
}

func runBuild(path string) error {
	// TODO(andydunstall): Build end to end (to output a binary).
	panic("not implemented")
}
