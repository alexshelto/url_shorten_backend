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


func UpdateUrlViewCount(url *models.Url) error {
        result := models.DB.Model(&url).Update("visit_count", url.VisitCount + 1).Error
        return result 
}

func UpdateUrlHash(url *models.Url, hash string) error {
    result := models.DB.Model(&url).Update("hashed", hash).Error
    return result
}

func GetUrlByHash(hash string) (error, models.Url) {
    queriedUrl := models.Url{}
    result := models.DB.First(&queriedUrl, hash).Error
    return result, queriedUrl
}

func FindByHashAndIncrementCount(hash string) (error, models.Url) {
    err, queriedUrl := GetUrlByHash(hash)
    if err != nil {
        return err, queriedUrl
    }
    fmt.Println(queriedUrl)

    err = UpdateUrlViewCount(&queriedUrl)
    if err != nil {
        return err, queriedUrl
    }
    return nil, queriedUrl
}


// POST: /api/v1/p
func CreateHashedPageV1(context *gin.Context) {
    body := CreateHashedUrlV1Body{}

    if err := context.BindJSON(&body); err!=nil {
        context.AbortWithError(http.StatusBadRequest,err)
        return
    }

    url := models.Url{Url: body.Url, Message: body.Message, ShowMsg: body.ShowMsgPage}
    result := models.DB.Create(&url)
    if result.Error != nil {
        context.AbortWithError(http.StatusInternalServerError, result.Error)
    }

    hash := BaseConversion.ConvertToBase62(url.Id)

    err := UpdateUrlHash(&url, hash)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        context.HTML(http.StatusNotFound, "404.html", nil)
        return
    }

    context.JSON(http.StatusAccepted,&body)
}


func GetPageFromHash(context *gin.Context) {
    hash := context.Param("hash")

    err, queriedUrl := FindByHashAndIncrementCount(hash)

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            fmt.Println("Not found")
            context.HTML(http.StatusNotFound, "404.html", nil)
            return
        }
        context.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    if queriedUrl.ShowMsg == true {
        context.HTML(http.StatusOK, "msg.html", gin.H{
            "message": queriedUrl.Message,
            "link": queriedUrl.Url,
            "visit_count": queriedUrl.VisitCount + 1,
        })
    } 
    context.Redirect(http.StatusPermanentRedirect, queriedUrl.Url)
}


func GetPageInfoFromHash(context *gin.Context) {
    hash := context.Param("hash")

    err, queriedUrl := FindByHashAndIncrementCount(hash)

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            fmt.Println("Not found")
            context.HTML(http.StatusNotFound, "404.html", nil)
            return
        }
        context.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    context.HTML(http.StatusOK, "msg.html", gin.H{
        "info": queriedUrl.Message,
        "link": queriedUrl.Url,
        "visit_count": queriedUrl.VisitCount + 1,
    })
}

