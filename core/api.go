package core

import (
	"fmt"
	"time"

	"github.com/imroc/req"
	"github.com/spf13/viper"
)

const (
	apiVersion = "v1"
	apiPath    = "api"
)

// ConnectLinkAPI encapsulates all API related code
type ConnectLinkAPI struct {
	ServerURL  string
	APIVersion string
	APIPath    string
	Username   string
	Password   string

	token      string
	validUntil time.Time
}

// NewConnectLinkAPI instance with the specified server and credentials
func NewConnectLinkAPI(serverURL, username, password string) *ConnectLinkAPI {
	fmt.Println("server:", viper.GetString("link.server_url"))
	return &ConnectLinkAPI{
		ServerURL:  serverURL,
		APIVersion: apiVersion,
		APIPath:    apiPath,
		Username:   username,
		Password:   password,
	}
}

// Ping the server and validate authentication is working properly
func (api *ConnectLinkAPI) Ping() (string, error) {
	api.RefreshToken()

	resp, err := get(api, "/ping")
	if err != nil {
		return "", err
	}

	result, _ := resp.ToString()
	return result, nil
}

// RefreshToken generates a new token with Connect Link when necessary
func (api *ConnectLinkAPI) RefreshToken() error {
	if time.Now().After(api.validUntil) {
		authResponse, err := api.Authenticate()
		if err != nil {
			er(err)
		}

		api.token = authResponse.Token
		api.validUntil = authResponse.ValidUntil
	}
	return nil
}

// Authenticate with Connect Link to retrieve a new token
func (api *ConnectLinkAPI) Authenticate() (*AuthenticationResponse, error) {
	resp, err := post(api, "/auth/request_token", &AuthenticationRequest{
		Username: api.Username,
		Password: api.Password,
	})

	if err != nil {
		return nil, err
	}

	result := &AuthenticationResponse{}
	resp.ToJSON(result)

	return result, nil
}

func get(api *ConnectLinkAPI, path string) (*req.Resp, error) {
	endpointURL := endPointURL(api.ServerURL, api.APIPath, api.APIVersion, path)
	header := apiHeader(api.token)

	fmt.Printf(" > GET: %s\n", endpointURL)
	return req.Get(endpointURL, header)
}

func post(api *ConnectLinkAPI, path string, any interface{}) (*req.Resp, error) {
	endpointURL := endPointURL(api.ServerURL, api.APIPath, api.APIVersion, path)
	header := apiHeader(api.token)

	fmt.Printf(" > POST: %s\n", endpointURL)
	return req.Post(endpointURL, header, req.BodyJSON(&any))
}

func endPointURL(serverURL, apiPath, apiVersion, path string) string {
	return fmt.Sprintf("%s/%s/%s%s", serverURL, apiPath, apiVersion, path)
}

func apiHeader(token string) req.Header {
	return req.Header{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"User-Agent":    "GoLink v0.0.0",
	}
}
