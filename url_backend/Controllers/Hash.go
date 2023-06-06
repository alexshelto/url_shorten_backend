package Controller

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


type CreateHashedPageV1Body struct {
  // json tag to de-serialize json body
   Url string `json:"url"`
}


/*
    POST: /api/v1/p
    BODY: { url: string }
*/
func CreateHashedPageV1(context *gin.Context) {
    body := CreateHashedPageV1Body{}

    if err := context.BindJSON(&body); err!=nil {
        context.AbortWithError(http.StatusBadRequest,err)
        return
    }

    fmt.Println(body)
    context.JSON(http.StatusAccepted,&body)
}


func GetPageFromHash(context *gin.Context) {
	context.JSON(http.StatusOK, "Shorten url")
}


