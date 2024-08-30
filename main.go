package main

import (
	"geemo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.BilibiliRoutersInit(r)

	r.Run(":8030")
}
