package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BilibiliRoutersInit(r *gin.Engine) {
	bilibiliRouters := r.Group("/bilibili")
	{
		bilibiliRouters.Any("/get-video-list", func(c *gin.Context) {
			fmt.Printf("get-video-list")
		})
	}
}
