package route

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gin_test/controler"
)





func InitApp() *gin.Engine {
    app := gin.New()


    app.Use(CORSMiddleware())



    app.Use(gin.Logger())
    app.Use(gin.Recovery())

    v1 := app.Group("/v1")
    {
        v1.GET("/index/test", func(cxt *gin.Context) {
            cxt.JSON(http.StatusOK, gin.H{"id_str": "222", "id_num": "2323"})
        })
    }


    index := app.Group("/index")
    {
        indexController := new(controler.IndexController)
        index.GET("/test", indexController.Test)
    }


    user := app.Group("/user")
    {
        userController := new(controler.UserController)
        user.POST("/test", userController.Test)
    }





    return app
}

