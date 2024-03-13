package wellknown

import (
	"fmt"
	"net/http"

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

	if query != fmt.Sprintf("acct:test@%s", c.Get("host").(string)) {
		return c.JSON(http.StatusBadRequest, nil)
	}

	resp := WebFingerResponse{
		Subject: fmt.Sprintf("acct:test@%s", c.Get("host").(string)),
		Links: []WebFingerLink{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: fmt.Sprintf("https://%s/users/test", c.Get("host").(string)),
			},
		},
	}
	return c.JSON(http.StatusOK, resp)
}
