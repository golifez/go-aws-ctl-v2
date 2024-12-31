package cmd

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golifez/go-aws-ctl-v2/internal/controller/s3"
)

// 启动WEB服务器
func StartWebServer(ctx context.Context, s *ghttp.Server) {
	s.SetPort(8888)
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")
	RootGroup(ctx, s) // 根组
	s.Run()
}

// /组
func RootGroup(ctx context.Context, s *ghttp.Server) {
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			s3.NewV1(),
		)
	})
}
