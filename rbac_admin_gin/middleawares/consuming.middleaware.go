package middleawares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Consuming func
func Consuming(c *gin.Context) {
	nowTime := time.Now()
	// 请求处理
	c.Next()
	//处理后获取消耗时间
	costTime := time.Since(nowTime)
	url := c.Request.URL.String()
	fmt.Printf("[time-consuming]the request URL: %s cost %v\n", url, costTime)
}
