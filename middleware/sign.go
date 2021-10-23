package middleware

import (
	"github.com/gin-gonic/gin"
	"goDemo/common"
	"goDemo/entity"
	"net/http"
)

func Sign() gin.HandlerFunc {

	return func(context *gin.Context) {

		res := entity.Result{}
		sign, err := common.VerifySign(context)
		if sign != nil {
			res.SetCode(entity.FAIL)
			res.SetMessage("Debug Sign")
			res.SetData(sign)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}
		if err != nil {
			res.SetCode(entity.FAIL)
			res.SetMessage(err.Error())
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}
		context.Next()
	}
}
