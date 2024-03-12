package wellknown

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type NodeInfoResponse struct {
	Links []NodeInfoLink `json:"links"`
}

type NodeInfoLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func NodeInfo(c echo.Context) error {
	host := fmt.Sprintf("%s:%s", os.Getenv("HOSTNAME"), os.Getenv("PORT"))
	resp := NodeInfoResponse{
		Links: []NodeInfoLink{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.1",
				Href: fmt.Sprintf("https://%s/nodeinfo/2.1", host),
			},
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: fmt.Sprintf("https://%s/nodeinfo/2.0", host),
			},
		},
	}
	return c.JSON(http.StatusOK, resp)
}
