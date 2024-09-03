package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var FakeHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Charset":  "UTF-8,*;q=0.5",
	"Accept-Language": "en-US,en;q=0.8",
	"Referer":         "https://www.bilibili.com",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
}

func BilibiliRoutersInit(r *gin.Engine) {
	bilibiliRouters := r.Group("/bilibili")
	{
		bilibiliRouters.GET("/get-video-list", GetVideoList)
		bilibiliRouters.GET("/get-season-list", GetSeasonList)
	}
}

func formatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	msg := ""
	if hours > 0 {
		msg = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, remainingSeconds)
	} else {
		msg = fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
	}
	return msg
}
