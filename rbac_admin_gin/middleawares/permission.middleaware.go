package middleawares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthPermissionMiddleaware() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()

		//请求处理
		c.Next()

		//处理后获取消耗时间
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}
