package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	_ "go-base/internal/logic"
	_ "go-base/internal/packed"

	"go-base/internal/cmd/apiserver"
)

func main() {
	apiserver.Main.Run(gctx.New())
}
