package http

import (
	clog "github.com/Chronostasys/centralog/centralog"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: nick
 * @Date: 2022/01/16 4:45 PM
 * @Version: 1.0.0
 */

func OK(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, data)
}
func Fail(ctx *gin.Context, err error) {
	defer clog.Sync()
	clog.Error(err.Error()).Log()
	if errorCode, ok := err.(*ErrorCode); ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorCode)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
}
