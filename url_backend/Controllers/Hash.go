package Controller

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


type CreateHashedPageV1Body struct {
    Url string `json:"url"`
    ShowMsgPage bool `json:"show_msg_page"`
    Message string `json:"message"`
}

/*
    POST: /api/v1/p
    BODY: { 
        url:             string (required)
        ShowMsgPage:     bool   (optional)
        Message:         string (optional)
    }
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
    // Reach out to db with hash..
    // if row has field of message page render that for them,
    // else redirect to the original 
	context.JSON(http.StatusOK, "Shorten url")
}


