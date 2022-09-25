package main

import (
	"github.com/rishabhgpt/fx-app/bundlefx"
	"github.com/rishabhgpt/fx-app/handler"
	"go.uber.org/fx"
)

func main() {
	fx.New(bundlefx.Module, fx.Provide(handler.New)).Run()
}
