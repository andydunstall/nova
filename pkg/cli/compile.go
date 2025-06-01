package cli

import (
	"fmt"
	"os"

	"github.com/andydunstall/nova/pkg/lex"
	"github.com/andydunstall/nova/pkg/print"
	"github.com/andydunstall/nova/pkg/syntax"
	"github.com/andydunstall/nova/pkg/types"
	"github.com/spf13/cobra"
)

func newCompileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compile path [flags]",
		Short: "compile a Nova program",
		Long:  `...`,
	}

	cmd.Run = func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			exitError(fmt.Errorf("compile: missing path"))
		}
		if len(args) > 1 {
			exitError(fmt.Errorf("compile: only one path is supported"))
		}

		if err := runCompile(args[0]); err != nil {
			exitError(fmt.Errorf("compile: %w", err))
		}
	}

	return cmd
}

func runCompile(path string) error {
	src, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read: %s: %w", path, err)
	}

	// Phase 1: Parse source into syntax AST.

	scanner := lex.NewScanner(src)
	syntaxAST, err := syntax.Parse(scanner)
	if err != nil {
		return fmt.Errorf("parse syntax: %w", err)
	}

	fmt.Println("syntax ast:")
	print.Print(syntaxAST)

	// Phase 2: Type checking.

	typeInfo, err := types.Check(syntaxAST)
	if err != nil {
		return fmt.Errorf("types: %w", err)
	}

	fmt.Println("type info:")
	print.Print(typeInfo)

	return nil
}
