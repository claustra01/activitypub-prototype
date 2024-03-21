package wellknown

import (
	"fmt"
	"net/http"

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
	resp := NodeInfoResponse{
		Links: []NodeInfoLink{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: fmt.Sprintf("https://%s/nodeinfo/2.0", c.Get("host").(string)),
			},
		},
	}
	return c.JSON(http.StatusOK, resp)
}
