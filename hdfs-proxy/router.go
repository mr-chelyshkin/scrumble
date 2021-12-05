package hdfs_proxy

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mr-chelyshkin/scrumble/hdfs-proxy/handlers"
	"github.com/mr-chelyshkin/scrumble/internal/sys"
)

type HdfsProxy struct {}

func (hp HdfsProxy) Echo(e *echo.Echo) {
	e.GET("/", handlers.Hello)
}

func (hp HdfsProxy) ThirdParty(e chan <- error) {
	ff := Config{}

	go func() {
		sys.ParseFileOnChange(
			e,
			"/Users/i.chelyshkin/Desktop/scrumble/_config/hdfs-proxy/other.toml",
			ff,
			ff.validate,
		)
	}()

}

func (hp HdfsProxy) Name() string {
	return "dhfs-proxy"
}

func (hp HdfsProxy) Utils(ctx context.Context) error {
	return nil
}
