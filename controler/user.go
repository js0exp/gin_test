package controler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type UserController struct{}


func (ctrl UserController) Test(cxt *gin.Context) {

    cxt.JSON(http.StatusOK, gin.H{"id_str": "222", "id_num": "2323"})
}

