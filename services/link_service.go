package services


import (
    "gorm.io/gorm"
    "alexshelto/url_shorten_service/models"
)


type LinkService struct {
    DB *gorm.DB
}

func NewLinkService(db *gorm.DB) *LinkService {
    return &LinkService{
        DB: db,
    }
}

func (ls *LinkService) CreateLink(link models.Link) error {
    err := ls.DB.Create(&link).Error
    return err
}

func (ls *LinkService) GetLinkById(linkId string) (models.Link, error) {
    var link models.Link
    err := ls.DB.First(&link, linkId).Error

    if err != nil {
        return models.Link{}, err
    }

    return link, nil
}
