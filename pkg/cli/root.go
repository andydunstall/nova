package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "nova [command] (flags)",
		SilenceUsage: true,
		Long: `Nova is a compiled statically typed programming language.

Build a Nova file (.nv extension) into an executable with:

  $ nova build ./prog.nv

Which will output the resulting binary to the same path as the input (with
the extension removed) by default.

See 'nova build -h' for the available options.

You can also invoke only the compiler (without the assembler or linker) to
output assembly:

  $ nova compile ./proc.nv

See 'nova compile -h' for the available options.
`,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cmd.AddCommand(
		newBuildCommand(),
		newCompileCommand(),
	)

	return cmd
}

func exitError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}

func init() {
	cobra.EnableCommandSorting = false
}
