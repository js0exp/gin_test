package main

import (
    "gin_test/route"
)


func main() {
    app := route.InitApp()

    app.Run(":8000") // listen and serve on 0.0.0.0:8080
}
