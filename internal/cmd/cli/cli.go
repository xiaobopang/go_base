package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"go-base/utility"
)

var (
	Main = gcmd.Command{
		Name:        "cli",
		Brief:       "A command-line tool demo",
		Description: "A command-line tool demo using GoFrame V2",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ver := parser.GetOpt("version")
			if ver != nil {
				utility.PrintVersionInfo()
				return
			}

			config := parser.GetOpt("config").String()
			if config != "" {
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config)
			}

			// json格式日志
			logFormat, err := g.Cfg().Get(ctx, "logger.cli.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 显示打印错误的文件行号
			g.Log("cli").SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)

			// 查看使用的配置文件是哪个
			configFile := g.Cfg().GetAdapter()
			g.Log("cli").Debugf(ctx, "use config file: %+v", configFile)

			// ****************** 以下部分为业务逻辑

			fmt.Printf("cli\n")

			g.Log("cli").Info(ctx, "foo")

			g.Log("cli").Error(ctx, errors.New("bar"))

			return
		},
	}
)
