package services

import (
	"alexshelto/url_shorten_service/models"
	"alexshelto/url_shorten_service/repositories"
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
    return ls.repo.Create(link)
}

func (ls *LinkService) GetLinkById(linkId uint) (models.Link, error) {
    return ls.repo.GetById(linkId)
}

func (ls *LinkService) GetLinkByShortenedUrl(shortenedUrl string) (models.Link, error) {
    return ls.repo.GetByShortenedUrl(shortenedUrl)
}
