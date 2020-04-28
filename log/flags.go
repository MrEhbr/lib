package log

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var Flags = []cli.Flag{
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "log.level",
			Aliases: []string{"log_level"},
			Value:   "info",
			EnvVars: []string{"LOG_LEVEL"},
		}),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "log.format",
			Aliases: []string{"log_format"},
			Value:   "json",
			EnvVars: []string{"LOG_FORMAT"},
		}),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "log.file",
			Aliases: []string{"log_file"},
			Value:   "stderr",
			EnvVars: []string{"LOG_FILE"},
		}),
}
