package repositories

import (
    "log"

    "gorm.io/gorm"
    "alexshelto/url_shorten_service/models"
)

type LinkRepository struct {
    DB *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
    return &LinkRepository{
        DB: db,
    }
}


func (repo *LinkRepository) Create(link models.Link) (models.Link, error) {
    if err := repo.DB.Create(&link).Error; err != nil {
        return models.Link{}, err
    }
    return link, nil
}


func (repo *LinkRepository) GetById(linkId uint) (models.Link, error) {
    var link models.Link

    if err := repo.DB.First(&link, linkId).Error; err != nil {
        return models.Link{}, err
    }

    // Immediately go and update the visit count 
    if err := repo.DB.Model(&link).Update("visit_count", link.VisitCount+1).Error; err != nil {
        return models.Link{}, err
    }

    return link, nil
}

func (repo *LinkRepository) UpdateShortenedUrlById(link models.Link, shortenedUrl string) (models.Link, error) {
    if err := repo.DB.Model(&link).Update("shortened_url", shortenedUrl).Error; err != nil {
        return link, err
    }
    return link, nil
}


func (repo *LinkRepository) GetByShortenedUrl(shortenedUrl string) (models.Link, error) {
    var link models.Link

    if err := repo.DB.First(&link, "shortened_url = ?", shortenedUrl).Error; err != nil {
        return models.Link{}, err
    }

    log.Println("Inside of get by shortenedURL")


    visitCount := link.VisitCount

    // Immediately go and update the visit count 
    if err := repo.DB.Model(&link).Update("visit_count", visitCount+1).Error; err != nil {
        return models.Link{}, err
    }

    return link, nil
}
