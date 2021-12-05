package hdfs_proxy

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mr-chelyshkin/scrumble/hdfs-proxy/handlers"
	"github.com/mr-chelyshkin/scrumble/internal/sys"
)

type HdfsProxy struct {}

func (hp HdfsProxy) Echo(e *echo.Echo) {
	e.GET("/", handlers.Hello)
}

func (hp HdfsProxy) ThirdParty(errCh chan error) {
	ff := Config{}
	fmt.Println("start")
	go func() {
		sys.ParseFileOnChange(
			errCh,
			"/Users/i.chelyshkin/Desktop/scrumble/_config/hdfs-proxy/other.toml",
			ff,
			ff.validate,
		)
	}()


}

func (hp HdfsProxy) Name() string {
	return "dhfs-proxy"
}
