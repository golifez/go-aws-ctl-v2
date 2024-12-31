package main

import (
	"github.com/golifez/go-aws-ctl-v2/internal/cmd"
	_ "github.com/golifez/go-aws-ctl-v2/internal/logic"
	_ "github.com/golifez/go-aws-ctl-v2/internal/packed"
	_ "github.com/golifez/go-aws-ctl-v2/src/cmd"
)

func main() {
	//cmd.Main.Run(gctx.GetInitCtx())
	cmd.Execute() // 执行工具行启动
}
