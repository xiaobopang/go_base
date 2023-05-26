package apiserver

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"go-base/router"

	"go-base/internal/service"
	"go-base/utility"
)

var (
	Main = gcmd.Command{
		Name:        "http-api",
		Brief:       "An API Server Demo",
		Description: "An API server demo using GoFrame V2",

		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 判断不带数据的选项是否存在时，可以通过GetOpt(name) != nil方式
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
			logFormat, err := g.Cfg().Get(ctx, "logger.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 异步打印日志 & 显示打印错误的文件行号, 对访问日志无效
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_STD | glog.F_FILE_LONG)

			configFile := g.Cfg().GetAdapter()
			g.Log().Debugf(ctx, "use config file: %+v", configFile)

			s := g.Server()

			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().TraceID,
					service.Middleware().ResponseHandler,
					service.Middleware().AccessUser,
				)

				router.Api(ctx, group) //加载路由

			})

			s.Run()
			return nil
		},
	}
)
