package v1

import (
	"github.com/gin-gonic/gin"
	"goDemo/entity"
	"net/http"
)

func AddMember(ctx *gin.Context) {

	res := entity.Result{}
	member := entity.Member{}
	if err := ctx.ShouldBind(&member); err != nil {
		res.SetCode(entity.FAIL)
		res.SetMessage(err.Error())
		ctx.JSON(http.StatusForbidden, res)
		ctx.Abort()
		return
	}
	data := map[string]interface{}{
		"name": member.Name,
		"age":  member.Age,
	}
	res.SetCode(entity.SUCCESS)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}
