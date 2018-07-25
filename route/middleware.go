package route

import (
    "github.com/gin-gonic/gin"
    "time"
    "log"
    "fmt"
)



func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            fmt.Println("OPTIONS")
            c.AbortWithStatus(200)
        } else {
            c.Next()
        }
    }
}



func LoggerTest() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        // 在gin上下文中定义变量
        c.Set("example", "12345")
        // 请求前
        c.Next()//处理请求
        // 请求后
        latency := time.Since(t)
        log.Print(t)
        log.Print(latency)
        fmt.Println("Hello, 世界")
        // access the status we are sending
        status := c.Writer.Status()
        log.Println(status)
    }
}