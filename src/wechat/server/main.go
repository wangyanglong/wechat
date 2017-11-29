package main

import (
	"net/http"
	"time"

	"wechat/view"

	"flag"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/jie123108/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	addr := "127.0.0.1:81"
	r := gin.Default()
	r.GET("/wx", view.VerifyServer)

	s := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 10,
	}
	err := gracehttp.Serve(s)
	if err != nil {
		glog.Errorf("wechat server run error:%s", err.Error())
	}
}
