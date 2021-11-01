package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"mmm/github-discussion/mutation"
	"net/http"
)

type SlackRequest struct {
	Token     string `json:"token" form:"token" query:"token"`
	Challenge string `json:"challenge" form:"challenge" query:"challenge"`
	Type      string `json:"type" form:"type" query:"type"`
}

func CreateGitHubDiscussions(c echo.Context) error {
	request := new(SlackRequest)
	if err := c.Bind(request); err != nil {
		log.Error(err)
		return c.NoContent(http.StatusOK)
	}

	err := mutation.CreateGitHubDiscussion()
	if err != nil {
		return err
	}

	response := map[string]string{"challenge": request.Challenge}
	if err := c.JSON(http.StatusOK, response); err != nil {
		return err
	}
	return nil
}
