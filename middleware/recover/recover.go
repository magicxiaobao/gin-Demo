package recover

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goDemo/common/alarm"
)

func Recover() gin.HandlerFunc {

	return func(context *gin.Context) {

		defer func() {
			if r := recover(); r != nil {
				alarm.Panic(fmt.Sprintf("%s", r))
			}
		}()
		context.Next()
	}
}
