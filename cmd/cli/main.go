package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	_ "go-base/internal/logic"
	_ "go-base/internal/packed"

	"go-base/internal/cmd/cli"
)

func main() {
	cli.Main.Run(gctx.New())
}
