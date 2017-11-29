package main

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/jie123108/glog"
)

func valiedServer(c *gin.Context) {

}

func main() {
	addr := "0.0.0.0:80"
	r := gin.Default()
	r.POST("/wx", valiedServer)

	s := new(http.Server)
	s.Addr = addr
	s.Handler = r
	s.ReadHeaderTimeout = time.Second * 10
	s.WriteTimeout = time.Second * 10
	err := gracehttp.Serve(s)
	if err != nil {
		glog.Errorf("wechat server run error:%s", err.Error())
	}
}
