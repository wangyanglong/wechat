package view

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jie123108/glog"
)

var myToken string = "my token"

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func VerifyServer(c *gin.Context) {
	glog.Infof("wechat VerifyServer")
	c.Request.ParseForm()
	data := new(VerifyDataReq)
	data.EchoStr = c.Request.Form["echostr"][0]
	data.Signature = c.Request.Form["signature"][0]
	data.Nonce = c.Request.Form["nonce"][0]
	data.Time = c.Request.Form["timestamp"][0]
	tmps := []string{myToken, data.Time, data.Nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	if str2sha1(tmpStr) == data.Signature {
		SetResp(c, 200, gin.H{})
		return
	} else {
		SetResp(c, 400, gin.H{})
		return
	}
}
