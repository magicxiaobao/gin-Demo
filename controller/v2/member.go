package v2

import (
	"github.com/gin-gonic/gin"
)

func AddMember(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"v2": "AddMember",
	})
}
