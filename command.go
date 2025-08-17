package command

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"

	yup "github.com/gloo-foo/framework"
)

type command yup.Inputs[string, flags]

func Perl(parameters ...any) yup.Command {
	return command(yup.Initialize[string, flags](parameters...))
}

func (p command) Executor() yup.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		// Get Perl code from positional arguments
		if len(p.Positional) == 0 {
			_, _ = fmt.Fprintf(stderr, "perl: no code specified\n")
			return fmt.Errorf("perl requires code to execute")
		}

		perlCode := strings.Join(p.Positional, " ")

		// Build perl command
		args := []string{}

		if bool(p.Flags.Loop) {
			args = append(args, "-n")
		}
		if bool(p.Flags.Print) {
			args = append(args, "-p")
		}
		if bool(p.Flags.AutoSplit) {
			args = append(args, "-a")
		}

		args = append(args, "-e", perlCode)

		// Execute perl
		cmd := exec.CommandContext(ctx, "perl", args...)
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		return cmd.Run()
	}
}
