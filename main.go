package main

import (
	_ "github.com/xframe-go/template/config"
	"github.com/xframe-go/x"
	"log/slog"
)

func main() {
	if err := x.App().Execute(); err != nil {
		slog.Error(err.Error())
	}
}
