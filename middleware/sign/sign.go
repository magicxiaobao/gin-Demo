package sign

import (
	"github.com/gin-gonic/gin"
	"github.com/magicxiaobao/ginDemo/common/function"
	"github.com/magicxiaobao/ginDemo/entity"
	"net/http"
)

func Sign() gin.HandlerFunc {

	return func(context *gin.Context) {

		res := entity.Result{}
		sign, err := function.VerifySign(context)
		if sign != nil {
			res = entity.ErrSignParam
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}
		if err != nil {
			res = entity.ErrSignParam
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}
		context.Next()
	}
}
