package user

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityPubPerson struct {
	Context           []string  `json:"@context"`
	Type              string    `json:"type"`
	Id                string    `json:"id"`
	Inbox             string    `json:"inbox"`
	Outbox            string    `json:"outbox"`
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

func generatePublicKeyPem() (string, error) {
	// RSA鍵ペアを生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	// 公開鍵をDER形式にエンコード
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode public key: %w", err)
	}

	// DER形式の公開鍵をPEM形式に変換
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(publicKeyPEM), nil
}

func MockUser(c echo.Context) error {

	url := fmt.Sprintf("https://%s/users/test", c.Get("host").(string))

	// 公開鍵を生成
	publicKeyPem, err := generatePublicKeyPem()
	if err != nil {
		return err
	}

	person := ActivityPubPerson{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		Type:              "Person",
		Id:                url,
		Inbox:             url + "/inbox",
		Outbox:            url + "/outbox",
		PrefferedUsername: "test",
		Name:              "Test User",
		Summary:           "The user in activitypub server made by claustra01",
		PublicKey: PublicKey{
			Type:         "Key",
			Id:           url + "#main-key",
			Owner:        url,
			PublicKeyPem: publicKeyPem,
		},
	}
	c.Response().Header().Set("Content-Type", "application/activity+json")
	return c.JSON(http.StatusOK, person)
}
