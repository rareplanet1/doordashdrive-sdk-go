package doordashdrive

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	baseURL = "https://openapi.doordash.com"
)

type Client struct {
	httpClient    *http.Client
	developerID   string
	keyID         string
	signingSecret string
}

func NewClient(developerID, keyID, signingSecret string) *Client {
	return &Client{
		httpClient:    &http.Client{Timeout: 10 * time.Second},
		developerID:   developerID,
		keyID:         keyID,
		signingSecret: signingSecret,
	}
}

func (c *Client) createJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "doordash",
		"iss": c.developerID,
		"kid": c.keyID,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"iat": time.Now().Unix(),
	})
	token.Header["dd-ver"] = "DD-JWT-V1"

	keyBytes, err := base64.URLEncoding.DecodeString(c.signingSecret)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(keyBytes)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (c *Client) doRequest(method, path string, body interface{}) (*http.Response, error) {
	resp, err := c.doRequestRaw(method, path, body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		apiError := &APIError{StatusCode: resp.StatusCode, Response: resp}

		if err := json.NewDecoder(resp.Body).Decode(apiError); err != nil {
			return nil, err
		}
		return nil, apiError
	}

	// jsonBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Print("Response: " + string(jsonBody) + "\n")

	return resp, nil
}

func (c *Client) doRequestRaw(method, path string, body interface{}) (*http.Response, error) {
	url := baseURL + path

	var bodyReader *bytes.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		// fmt.Print("Request: " + string(jsonBody) + "\n")
		bodyReader = bytes.NewReader(jsonBody)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	token, err := c.createJWT()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}
