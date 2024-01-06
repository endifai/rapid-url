package controllers

import (
	"errors"
	"net/http"
	"rapidurl/internal/repositories"

	"github.com/labstack/echo/v4"
)

type LinkController struct {
	linkRepository *repositories.LinkRepository
}

func NewLinkController(repo *repositories.LinkRepository) *LinkController {
	return &LinkController{
		linkRepository: repo,
	}
}

type CreateLinkBody struct {
	Link string `json:"link"`
}

func (controller *LinkController) CreateLink(c echo.Context) error {
	body := new(CreateLinkBody)

	if err := c.Bind(body); err != nil {
		return err
	}

	createdLink, error := controller.linkRepository.CreateLink(body.Link)

	if error != nil {
		return error
	}

	return c.JSON(http.StatusCreated, createdLink)
}

func (controller *LinkController) GetLink(c echo.Context) error {
	shortLink := c.Param("hash")

	entity, err := controller.linkRepository.GetByShortLink(shortLink)

	if err != nil {
		switch {
		case errors.Is(err, repositories.ErrNotFound):
			return c.String(http.StatusNotFound, "Not found")
		}

		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, entity.Link)
}
