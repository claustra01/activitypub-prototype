package wellknown

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type WebFingerResponse struct {
	Subject string          `json:"subject"`
	Links   []WebFingerLink `json:"links"`
}

type WebFingerLink struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

func WebFinger(c echo.Context) error {
	query := c.QueryParam("resource")
	host := fmt.Sprintf("%s:%s", os.Getenv("HOSTNAME"), os.Getenv("PORT"))

	if query != fmt.Sprintf("acct:test@%s", host) {
		return c.JSON(http.StatusBadRequest, nil)
	}

	resp := WebFingerResponse{
		Subject: fmt.Sprintf("acct:test@%s", host),
		Links: []WebFingerLink{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: fmt.Sprintf("https://%s/users/test", host),
			},
		},
	}
	return c.JSON(http.StatusOK, resp)
}
