package utility

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gbuild"
)

// PrintVersionInfo for application
func PrintVersionInfo() {
	info := gbuild.Info()
	if info.Git == "" {
		info.Git = "none"
	}

	fmt.Printf(`App Version: %s
Git Commit:  %s
Build Time:  %s
Go Version:  %s
GF Version:  %s
`, info.Data["version"], // gfcli.build.varMap.version in hack/config.yaml
		info.Git,
		info.Time,
		info.Golang,
		info.GoFrame,
	)
}
