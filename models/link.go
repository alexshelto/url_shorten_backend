package models

import (
    "gorm.io/gorm"
)

type Link struct {
    gorm.Model

    OriginalUrl  string 
    ShortenedUrl string 
    VisitCount   int 
}

type CreateLinkRequest struct {
    OriginalUrl string `json:"original_url" binding:"required"`
}


type CreateLinkFormData struct {
    OriginalUrl string `form:"url" binding:"required"`
}
