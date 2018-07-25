package controler

import (
    "github.com/gin-gonic/gin"
    "strconv"
    "fmt"
    "net/http"
)

type IndexController struct{}


func (ctrl IndexController) Test(cxt *gin.Context) {
    name := cxt.Query("name")
    id := cxt.DefaultQuery("id", "222")
    idNum, _ := strconv.Atoi(id)
    idStr := strconv.Itoa(idNum)
    fmt.Println(name)

    username := cxt.PostForm("username")
    password := cxt.DefaultPostForm("password", "343434")
    fmt.Println(username + "--------------" + password)

    var user map[string]interface{}
    //body, _ := ioutil.ReadAll(cxt.Request.Body)
    //json.Unmarshal(body, &user)

    fmt.Println("user_image_url----post_json----")
    cxt.BindJSON(&user)
    fmt.Println(user)
    cxt.JSON(http.StatusOK, gin.H{"id_str": idStr, "id_num": idNum, "name": name})
}

func (ctrl IndexController) Login(cxt *gin.Context)  {
    cxt.JSON(http.StatusOK, gin.H{"user": "test", "Message": "test", "Number": 123})
}


func (ctrl IndexController) Logout(cxt *gin.Context)  {
    cxt.JSON(http.StatusOK, gin.H{"user": "test", "Message": "test", "Number": 123})
}