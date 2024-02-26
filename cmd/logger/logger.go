package logger

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

// InitLogger initializes the root logger with the command line flags.
func InitLogger(c *cli.Context) {
	var (
		slogVerbosity = log.FromLegacyLevel(c.Int(flags.Verbosity.Name))
	)

	if c.Bool(flags.LogJSON.Name) {
		glogger := log.NewGlogHandler(log.NewGlogHandler(log.JSONHandler(os.Stdout)))
		glogger.Verbosity(slogVerbosity)
		log.SetDefault(log.NewLogger(glogger))
	} else {
		glogger := log.NewGlogHandler(log.NewTerminalHandler(os.Stdout, true))
		glogger.Verbosity(slogVerbosity)
		log.SetDefault(log.NewLogger(glogger))
	}
}
