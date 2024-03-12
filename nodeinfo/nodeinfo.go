package nodeinfo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type NodeInfoResponse struct {
	SoftWare          SoftWare `json:"software"`
	Version           string   `json:"version"`
	Protocols         []string `json:"protocols"`
	OpenRegistrations bool     `json:"openRegistrations"`
	Usage             Usage    `json:"usage"`
}

type SoftWare struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Usage struct {
	Users UsageUsers `json:"users"`
}

type UsageUsers struct {
	Total int `json:"total"`
}

func NodeInfo2_0(e echo.Context) error {
	resp := NodeInfoResponse{
		SoftWare: SoftWare{
			Name:    "activitypub-prototype",
			Version: "0.0.1",
		},
		Version:           "2.0",
		Protocols:         []string{"activitypub"},
		OpenRegistrations: false,
		Usage: Usage{
			Users: UsageUsers{
				Total: 1,
			},
		},
	}
	return e.JSON(http.StatusOK, resp)
}

func NodeInfo2_1(e echo.Context) error {
	resp := NodeInfoResponse{
		SoftWare: SoftWare{
			Name:    "activitypub-prototype",
			Version: "0.0.1",
		},
		Version:           "2.1",
		Protocols:         []string{"activitypub"},
		OpenRegistrations: false,
		Usage: Usage{
			Users: UsageUsers{
				Total: 1,
			},
		},
	}
	return e.JSON(http.StatusOK, resp)
}
