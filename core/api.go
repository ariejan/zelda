package core

import (
	"fmt"
	"time"

	"github.com/imroc/req"
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

// FetchInstallations fetches an up-to-date list of installations
func (api *ConnectLinkAPI) FetchInstallations() (string, error) {
	api.RefreshToken()

	resp, err := get(api, "/installations")
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{200})

	result, _ := resp.ToString()
	return result, nil
}

// FetchZones fetches an up-to-date list of zones for the specified installationID
func (api *ConnectLinkAPI) FetchZones(installationID int) (string, error) {
	api.RefreshToken()

	resp, err := get(api, fmt.Sprintf("/installations/%d/zones", installationID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{200})

	result, _ := resp.ToString()
	return result, nil
}

// FetchElements fetches an up-to-date list of elements for the specified installationID and zoneID
func (api *ConnectLinkAPI) FetchElements(installationID, zoneID int) (string, error) {
	api.RefreshToken()

	resp, err := get(api, fmt.Sprintf("/installations/%d/zones/%d/elements", installationID, zoneID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{200})

	result, _ := resp.ToString()
	return result, nil
}

// FetchAlerts fetches an up-to-date list of alerts for the specified installationID
func (api *ConnectLinkAPI) FetchAlerts(installationID int) (string, error) {
	api.RefreshToken()

	resp, err := get(api, fmt.Sprintf("/installations/%d/alerts", installationID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{200})

	result, _ := resp.ToString()
	return result, nil
}

// FetchWebhooks fetches an up-to-date list of your webhooks for the installation
func (api *ConnectLinkAPI) FetchWebhooks(installationID int) (string, error) {
	api.RefreshToken()

	resp, err := get(api, fmt.Sprintf("/installations/%d/webhooks", installationID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{200})

	result, _ := resp.ToString()
	return result, nil
}

// FetchWebhook fetches info about a specific webhook
func (api *ConnectLinkAPI) FetchWebhook(installationID, webhookID int) (string, error) {
	api.RefreshToken()

	resp, err := get(api, fmt.Sprintf("/installations/%d/webhooks/%d", installationID, webhookID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{200})

	result, _ := resp.ToString()
	return result, nil
}

// CreateWebhook creates a new webhook
func (api *ConnectLinkAPI) CreateWebhook(installationID int, endpoint, token string) (string, error) {
	api.RefreshToken()

	resp, err := post(api, fmt.Sprintf("/installations/%d/webhooks", installationID), &WebhookRequest{
		Endpoint: endpoint,
	})
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{201})

	result, _ := resp.ToString()
	return result, nil
}

// DeleteWebhook deletes the specified webhook
func (api *ConnectLinkAPI) DeleteWebhook(installationID, webhookID int) (string, error) {
	api.RefreshToken()

	resp, err := delete(api, fmt.Sprintf("/installations/%d/webhooks/%d", installationID, webhookID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{204})

	result, _ := resp.ToString()
	return result, nil
}

// TestWebhook requests a test event to be fired for this webhook
func (api *ConnectLinkAPI) TestWebhook(installationID, webhookID int) (string, error) {
	api.RefreshToken()

	resp, err := postWithoutBody(api, fmt.Sprintf("/installations/%d/webhooks/%d/test", installationID, webhookID))
	if err != nil {
		return "", err
	}

	validateStatusCode(resp, []int{201})

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

	fmt.Printf("~~> GET: %s\n", endpointURL)
	return req.Get(endpointURL, header)
}

func post(api *ConnectLinkAPI, path string, any interface{}) (*req.Resp, error) {
	endpointURL := endPointURL(api.ServerURL, api.APIPath, api.APIVersion, path)
	header := apiHeader(api.token)

	fmt.Printf("~~> POST: %s\n", endpointURL)
	return req.Post(endpointURL, header, req.BodyJSON(&any))
}

func postWithoutBody(api *ConnectLinkAPI, path string) (*req.Resp, error) {
	endpointURL := endPointURL(api.ServerURL, api.APIPath, api.APIVersion, path)
	header := apiHeader(api.token)

	fmt.Printf("~~> POST: %s\n", endpointURL)
	return req.Post(endpointURL, header)
}

func delete(api *ConnectLinkAPI, path string) (*req.Resp, error) {
	endpointURL := endPointURL(api.ServerURL, api.APIPath, api.APIVersion, path)
	header := apiHeader(api.token)

	fmt.Printf("~~> DELETE: %s\n", endpointURL)
	return req.Delete(endpointURL, header)
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
