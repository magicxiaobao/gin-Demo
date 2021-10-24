package v1

import (
	"github.com/gin-gonic/gin"
	"goDemo/common/alarm"
	"goDemo/entity"
	"net/http"
)

func AddMember(ctx *gin.Context) {

	res := entity.Result{}
	member := entity.Member{}
	if err := ctx.ShouldBind(&member); err != nil {
		res = entity.ErrParam
		ctx.JSON(http.StatusForbidden, res)
		alarm.Error(err.Error())
		ctx.Abort()
		return
	}
	if member.Name == "xiaobao" {
		panic("xiaobao is not allowed")
	}
	data := map[string]interface{}{
		"name": member.Name,
		"age":  member.Age,
	}
	res = entity.OK.WithData(data)
	ctx.JSON(http.StatusOK, res)
}
