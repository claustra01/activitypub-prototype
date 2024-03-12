package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityPubPerson struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}

func MockUser(c echo.Context) error {
	person := ActivityPubPerson{
		Context: "https://www.w3.org/ns/activitystreams",
		Type:    "Person",
		Name:    "test",
	}
	return c.JSON(http.StatusOK, person)
}
