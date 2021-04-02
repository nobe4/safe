package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nobe4/safe/internal/entropy"
	"github.com/nobe4/safe/internal/logger"
	"github.com/nobe4/safe/internal/shell"
	"github.com/nobe4/safe/internal/writer"
)

var (
	Version string
	Commit  string
	Build   string
)

func main() {
	// Setup the flags
	regexp := flag.String("regexp", "", "a regexp to hide")
	verbosity := flag.Int("verbosity", 2, "verbosity level (0: nothing, 1: errors, 2: warnings, 3: info, 4: debug)")
	censor := flag.String("censor", "X", "censor character to use")
	dict := flag.String("dict", "ascii", "dictionnary for entropy filtering ("+entropy.List()+")")
	threshold := flag.Float64("threshold", 3.0, "threshold to apply filtering (debug with verbosity 1)")
	version := flag.Bool("version", false, "show version information")
	flag.Parse()

	if version != nil && *version {
		fmt.Printf("safe: version %s, commit %s, build %s\n", Version, Commit, Build)
		return
	}

	logger.Debug("Set logger level to:", *verbosity)
	logger.SetLevel(*verbosity)

	safeWriter, err := writer.New(regexp, censor, dict, threshold)
	if err != nil {
		logger.Error(err)
	}

	useShell, err := shell.IsShell()
	if err != nil {
		logger.Error(err)
	}

	if err := run(safeWriter, useShell); err != nil {
		logger.Error(err)
	}
}

func run(out io.Writer, useShell bool) error {
	if useShell {
		logger.Info("Use a shell.")
		return shell.Start(out)
	}

	logger.Info("Use a stdin pipe.")

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()

		logger.Debug("Read a line:", line)

		if _, err := out.Write([]byte(line + "\n")); err != nil {
			return err
		}
	}

	return input.Err()
}
