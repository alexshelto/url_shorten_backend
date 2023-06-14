package services

import (
	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/repositories"
	"alexshelto/url_shorten_service/utils"
)


type LinkService struct {
    repo *repositories.LinkRepository
}

func NewLinkService(repo *repositories.LinkRepository) *LinkService {
    return &LinkService{
        repo: repo,
    }
}

func (ls *LinkService) CreateLink(link models.Link) (models.Link, error) {
    link, err := ls.repo.Create(link)
    if err != nil {
        return link, err
    }

    generatedUrl := BaseConversion.ConvertToBase62(link.ID)
    return ls.repo.UpdateShortenedUrlById(link, generatedUrl)
}

func (ls *LinkService) GetLinkById(linkId uint) (models.Link, error) {
    return ls.repo.GetById(linkId)
}

func (ls *LinkService) GetLinkByShortenedUrl(shortenedUrl string) (models.Link, error) {
    return ls.repo.GetByShortenedUrl(shortenedUrl)
}
