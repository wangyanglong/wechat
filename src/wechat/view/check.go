package view

import (
	"encoding/json"
	"io/ioutil"

	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
	"github.com/jie123108/glog"
	valid "gopkg.in/go-playground/validator.v8"
)

func SetResp(c *gin.Context, status int, jso gin.H) {
	seted := context.Get(c.Request, "status_code")
	if seted == nil {
		c.JSON(status, jso)
		context.Set(c.Request, "status_code", true)
	} else {
		glog.Infof("find a response body, ignore current Resp data: status: %d, json: %v", status, jso)
	}
}

func check_request_body(c *gin.Context, body []byte, obj interface{}) bool {
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	decoder.UseNumber()
	if err := decoder.Decode(obj); err != nil {
		glog.Errorf("Invalid Body [[%s]], err: %v", string(body), err)
		return false
	}

	config := &valid.Config{TagName: "validate"}
	validator := valid.New(config)
	if err := validator.Struct(obj); err != nil {
		glog.Errorf("Invalid Input Json err: %s,[%s]! ", err.Error(), string(body))
		return false
	}
	return true
}

func CheckRequestBody(c *gin.Context, obj interface{}) bool {
	body, _ := ioutil.ReadAll(c.Request.Body)
	ret := check_request_body(c, body, obj)
	if !ret {
		glog.Errorf("check request body error:%s", string(body))
	}
	return ret
}

type VerifyDataReq struct {
	Signature string `json:"signature"`
	Time      string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	EchoStr   string `json:"echostr"`
}
