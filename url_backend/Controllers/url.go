package Controller

import (
	"alexshelto/url_shorten_service/Models"
	"alexshelto/url_shorten_service/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


/*
The POST body data to create a url 
*/
type CreateHashedUrlV1Body struct {
    Url string `json:"url"`
    ShowMsgPage bool `json:"show_msg_page"`
    Message string `json:"message"`
}



// POST: /api/v1/p
func CreateHashedPageV1(context *gin.Context) {
    body := CreateHashedUrlV1Body{}

    if err := context.BindJSON(&body); err!=nil {
        context.AbortWithError(http.StatusBadRequest,err)
        return
    }

    fmt.Println(body)

    url := models.Url{Url: body.Url, Message: body.Message, ShowMsg: body.ShowMsgPage}
    result := models.DB.Create(&url)

    if result.Error != nil {
        context.AbortWithError(http.StatusInternalServerError, result.Error)
    }

    hash := BaseConversion.ConvertToBase62(url.Id)
    fmt.Println("item had id: ", url.Id, " created hash: ", hash)
    url.Hashed = hash

    result = models.DB.Model(&url).Update("hashed", hash)

    if result.Error != nil {
        context.AbortWithError(http.StatusInternalServerError, result.Error)
    }

    context.JSON(http.StatusAccepted,&body)
}


func GetPageFromHash(context *gin.Context) {
    // Reach out to db with hash..
    // if row has field of message page render that for them,
    // else redirect to the original 
	context.JSON(http.StatusOK, "Shorten url")
}


