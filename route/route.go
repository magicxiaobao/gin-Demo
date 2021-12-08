package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/magicxiaobao/ginDemo/common/function"
	"github.com/magicxiaobao/ginDemo/controller/v1"
	v2 "github.com/magicxiaobao/ginDemo/controller/v2"
	"github.com/magicxiaobao/ginDemo/middleware/sign"
	"github.com/magicxiaobao/ginDemo/validator"
	"net/url"
	"strconv"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/sn", SignDemo)
	// v1
	group1 := engine.Group("v1")
	{
		group1.Any("/member/add", v1.AddMember)
	}
	group2 := engine.Group("v2")
	group2.Use(sign.Sign())
	{
		group2.Any("/member/add", v2.AddMember)
	}
	// 绑定验证器-v10
	binding.Validator = new(validator.DefaultValidator)
}

func SignDemo(ctx *gin.Context) {

	ts := strconv.FormatInt(function.GetCurrentTimeStamp(), 10)
	response := map[string]interface{}{}
	params := url.Values{
		"name": []string{"guoguo"},
		"age":  []string{"9"},
		"ts":   []string{ts},
	}
	response["sn"] = function.CreateSign(params)
	response["ts"] = ts
	function.RetJson("200", "", response, ctx)
}
