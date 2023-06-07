package Controller

import (
	"alexshelto/url_shorten_service/Models"
	"alexshelto/url_shorten_service/Utils"
	"fmt"
	"net/http"

    "errors"
    "gorm.io/gorm"

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
    hash := context.Param("hash")

    queriedUrl := models.Url{}

    err := models.DB.First(&queriedUrl, hash).Error
    fmt.Println(err)

    if errors.Is(err, gorm.ErrRecordNotFound) {
        fmt.Println("Not found")
        context.HTML(http.StatusNotFound, "404.html", nil)
        return
    }

    if err != nil {
        context.AbortWithError(http.StatusInternalServerError, err)
    }

    if queriedUrl.ShowMsg == true {
        context.HTML(http.StatusOK, "msg.html", gin.H{
            "message": queriedUrl.Message,
            "link": queriedUrl.Url,
        })
    } 

    context.Redirect(http.StatusPermanentRedirect, queriedUrl.Url)
}



