package main

import (
	_ "github.com/golifez/go-aws-ctl-v2/cmd"
	_ "github.com/golifez/go-aws-ctl-v2/internal/logic"

	"github.com/golifez/go-aws-ctl-v2/internal/cmd"
	_ "github.com/golifez/go-aws-ctl-v2/internal/packed"
)

func main() {
	//cmd.Main.Run(gctx.GetInitCtx())
	cmd.Execute() // 执行工具行启动
}
