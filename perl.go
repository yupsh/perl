package perl

import (
	"context"
	"fmt"
	"io"

	execCmd "github.com/yupsh/exec"
	execOpt "github.com/yupsh/exec/opt"
	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/perl/opt"
)

// Flags represents the configuration options for perl commands
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Perl creates a new perl command with the given parameters
func Perl(parameters ...any) yup.Command {
	return command(opt.Args[string, Flags](parameters...))
}

func (c command) Execute(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	// Build perl command arguments
	args := make([]any, 0, len(c.Positional)+20)
	args = append(args, "perl")

	// Add flag-based arguments first
	args = append(args, c.buildPerlFlags()...)

	// Add positional arguments
	for _, arg := range c.Positional {
		args = append(args, any(arg))
	}

	// Convert to exec options
	execOpts := c.buildExecOptions()
	args = append(args, execOpts...)

	// Create and execute the underlying exec command
	execCommand := execCmd.Exec(args...)
	return execCommand.Execute(ctx, stdin, stdout, stderr)
}

func (c command) buildPerlFlags() []any {
	var flags []any

	// Add perl-specific flags
	if bool(c.Flags.Warnings) {
		flags = append(flags, "-w")
	}
	if bool(c.Flags.Taint) {
		flags = append(flags, "-T")
	}
	if bool(c.Flags.Debug) {
		flags = append(flags, "-d")
	}
	if bool(c.Flags.CheckSyntax) {
		flags = append(flags, "-c")
	}
	if bool(c.Flags.InPlace) {
		flags = append(flags, "-i")
	}
	if bool(c.Flags.Print) {
		flags = append(flags, "-p")
	}
	if bool(c.Flags.Loop) {
		flags = append(flags, "-n")
	}
	if bool(c.Flags.AutoSplit) {
		flags = append(flags, "-a")
	}

	// Add module includes
	if c.Flags.Module != "" {
		flags = append(flags, "-M", string(c.Flags.Module))
	}

	// Add library paths
	if c.Flags.LibPath != "" {
		flags = append(flags, "-I", string(c.Flags.LibPath))
	}

	// Add script if specified
	if c.Flags.Script != "" {
		flags = append(flags, "-e", string(c.Flags.Script))
	}

	return flags
}

func (c command) buildExecOptions() []any {
	var opts []any

	// Always inherit environment for perl
	opts = append(opts, execOpt.InheritEnv)

	return opts
}

func (c command) String() string {
	return fmt.Sprintf("perl %v", c.Positional)
}

// Wrapper commands for common perl operations

// Execute creates a perl -e command (execute script)
func Execute(script string, parameters ...any) yup.Command {
	args := make([]any, 0, 2+len(parameters))
	args = append(args, "-e", script)
	args = append(args, parameters...)
	return Perl(args...)
}

// InPlaceEdit creates a perl -pi -e command (in-place edit)
func InPlaceEdit(script string, parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	cmd.Flags.InPlace = localopt.InPlace
	cmd.Flags.Print = localopt.Print
	cmd.Flags.Script = localopt.Script(script)
	return cmd
}

// Loop creates a perl -n -e command (loop without printing)
func Loop(script string, parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	cmd.Flags.Loop = localopt.Loop
	cmd.Flags.Script = localopt.Script(script)
	return cmd
}

// PrintLoop creates a perl -p -e command (loop with printing)
func PrintLoop(script string, parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	cmd.Flags.Print = localopt.Print
	cmd.Flags.Script = localopt.Script(script)
	return cmd
}

// AutoSplitLoop creates a perl -na -e command (auto-split loop)
func AutoSplitLoop(script string, parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	cmd.Flags.Loop = localopt.Loop
	cmd.Flags.AutoSplit = localopt.AutoSplit
	cmd.Flags.Script = localopt.Script(script)
	return cmd
}

// AutoSplitPrint creates a perl -pa -e command (auto-split with print)
func AutoSplitPrint(script string, parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	cmd.Flags.Print = localopt.Print
	cmd.Flags.AutoSplit = localopt.AutoSplit
	cmd.Flags.Script = localopt.Script(script)
	return cmd
}

// CheckSyntax creates a perl -c command (syntax check)
func CheckSyntax(parameters ...any) yup.Command {
	args := append([]any{"-c"}, parameters...)
	return Perl(args...)
}

// WithWarnings creates a perl -w command (with warnings)
func WithWarnings(parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	cmd.Flags.Warnings = localopt.Warnings
	return cmd
}

// WithModule creates a perl -M command (with module)
func WithModule(module string, parameters ...any) yup.Command {
	args := make([]any, 0, 2+len(parameters))
	args = append(args, "-M", module)
	args = append(args, parameters...)
	return Perl(args...)
}

// One-liner helpers for common patterns

// Substitute creates a perl -pi -e 's/old/new/g' command
func Substitute(pattern, replacement string, parameters ...any) yup.Command {
	script := fmt.Sprintf("s/%s/%s/g", pattern, replacement)
	return InPlaceEdit(script, parameters...)
}

// GlobalSubstitute creates a perl -pi -e 's/old/new/g' command (alias for Substitute)
func GlobalSubstitute(pattern, replacement string, parameters ...any) yup.Command {
	return Substitute(pattern, replacement, parameters...)
}

// Print creates a perl -ne 'print if /pattern/' command
func Print(pattern string, parameters ...any) yup.Command {
	script := fmt.Sprintf("print if /%s/", pattern)
	return Loop(script, parameters...)
}

// Count creates a perl -ne 'END{print $.}' command (count lines)
func Count(parameters ...any) yup.Command {
	script := "END{print $.}"
	return Loop(script, parameters...)
}

// Sum creates a perl -ne '$sum += $_; END{print $sum}' command
func Sum(parameters ...any) yup.Command {
	script := "$sum += $_; END{print $sum}"
	return Loop(script, parameters...)
}
