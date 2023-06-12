package models

import (
    "gorm.io/gorm"
)

type Link struct {
    gorm.Model

    OriginalUrl string `gorm:"column:original_url"`
    ShortenedUrl string `gorm:"column:shortened_url"`
    VisitCount int `gorm:"column:visit_count"`
}

