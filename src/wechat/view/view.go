package view

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jie123108/glog"
)

var myToken string = "myToken"

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func VerifyServer(c *gin.Context) {
	glog.Infof("wechat VerifyServer")
	c.Request.ParseForm()
	data := new(VerifyDataReq)
	data.EchoStr = CheckGetParam(c, "echostr")
	data.Signature = CheckGetParam(c, "signature")
	data.Nonce = CheckGetParam(c, "nonce")
	data.Time = CheckGetParam(c, "timestamp")
	tmps := []string{myToken, data.Time, data.Nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	if str2sha1(tmpStr) == data.Signature {
		c.String(200, data.EchoStr)
		return
	} else {
		SetResp(c, 400, gin.H{})
		return
	}
}
