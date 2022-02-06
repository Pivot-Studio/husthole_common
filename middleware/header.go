package middleware

import (
	"github.com/Pivot-Studio/husthole_common/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	AppIDFieldName = "x-appid"
	AppIDStart     = AppIDWx
	AppIDWx        = 103701
	AppIDH5        = 103702
	AppIDAndroid   = 103703
	AppIDIOS       = 103704
	AppIDEnd       = AppIDIOS
)

func CheckHeader(context *gin.Context) {
	CheckAppID(context)
	context.Next()
}

func CheckAppID(context *gin.Context) {
	defer log.Sync()
	appidStr := context.GetHeader(AppIDFieldName)
	appid, err := strconv.Atoi(appidStr)
	if err != nil || appid > AppIDEnd || appid < AppIDStart {
		log.CtxError(context, "invalid appid")
		context.Abort()
	}
}
