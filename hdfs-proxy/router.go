package hdfs_proxy

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mr-chelyshkin/scrumble/hdfs-proxy/handlers"
	"github.com/mr-chelyshkin/scrumble/internal/sys"
	"github.com/pkg/errors"
)

type HdfsProxy struct {}

func (hp HdfsProxy) Echo(e *echo.Echo) {
	e.GET("/", handlers.Hello)
}

func (hp HdfsProxy) ThirdParty(ctx context.Context) error {
	ff := Config{}
	err := sys.ParseFileOnChange(
		ctx,
		"/Users/i.chelyshkin/Desktop/scrumble/_config/hdfs-proxy/other.toml",
		ff,
		ff.validate,
	)
	if err != nil {
		fmt.Println(err)
	}

	return errors.Errorf("asd")
}

func (hp HdfsProxy) Name() string {
	return "dhfs-proxy"
}
