package config

import (
	"github.com/xframe-go/x"
	"github.com/xframe-go/x/contracts"
	"github.com/xframe-go/x/utils/env"
	"github.com/xframe-go/x/xlog"
	"os"
)

func init() {
	x.App().Config().Register(&xlog.Config{
		Default: env.String("LOG_CHANNEL", "stdout"),

		Channels: map[string]contracts.Logger{

			"discard": xlog.NewDiscardLogger(),

			"single": xlog.NewFileLogger(xlog.FileLoggerConfig{
				Format: env.String("LOG_FORMAT", xlog.TextFormatter),

				Writer: xlog.NewFileWriter("x", xlog.WithRotate(xlog.SingleRotate)),
			}),

			"daily": xlog.NewFileLogger(xlog.FileLoggerConfig{
				Format: env.String("LOG_FORMAT", xlog.TextFormatter),

				Writer: xlog.NewFileWriter("x", xlog.WithRotate(xlog.DailyRotate)),
			}),

			"stdout": xlog.NewFileLogger(xlog.FileLoggerConfig{
				Format: env.String("LOG_FORMAT", xlog.TextFormatter),

				Writer: os.Stdout,
			}),
		},
	})
}
