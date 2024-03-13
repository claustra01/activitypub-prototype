package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityPubPerson struct {
	Context           []string  `json:"@context"`
	Type              string    `json:"type"`
	Id                string    `json:"id"`
	Inbox             string    `json:"inbox"`
	PrefferedUsername string    `json:"prefferedUsername"`
	Name              string    `json:"name"`
	Summary           string    `json:"summary"`
	PublicKey         PublicKey `json:"publicKey"`
}

type PublicKey struct {
	Type         string `json:"type"`
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

func MockUser(c echo.Context) error {

	url := fmt.Sprintf("https://%s/users/test", c.Get("host").(string))
	person := ActivityPubPerson{
		Context:           []string{"https://www.w3.org/ns/activitystreams"},
		Type:              "Person",
		Id:                url,
		Inbox:             url + "/inbox",
		PrefferedUsername: "test",
		Name:              "Test User",
		Summary:           "The user in activitypub server made by claustra01",
		PublicKey: PublicKey{
			Type:         "Key",
			Id:           url + "#main-key",
			Owner:        url,
			PublicKeyPem: "key is here",
		},
	}
	c.Response().Header().Set("Content-Type", "application/activity+json")
	return c.JSON(http.StatusOK, person)
}
