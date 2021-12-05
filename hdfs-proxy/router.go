package hdfs_proxy

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mr-chelyshkin/scrumble/hdfs-proxy/handlers"
	"github.com/mr-chelyshkin/scrumble/internal/sys"
)

func Route(e *echo.Echo) {
	e.GET("/", handlers.Hello)
}

func In(ctx context.Context)  {
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
}
