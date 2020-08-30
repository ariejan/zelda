package core

import "time"

// PingResponse containing, hopefully, a PONG
type PingResponse struct {
	Response string
}

// AuthenticationRequest is a request to authenticate
type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthenticationResponse is the response to our AuthenticationRequest
type AuthenticationResponse struct {
	Token      string
	ValidUntil time.Time
}

// CustomerOrganisation represents a customer organisation for an instalaltion
type CustomerOrganisation struct {
	ID   int
	Name string
}

// ServiceOrganisation represents the service organisation for an installation
type ServiceOrganisation struct {
	ID   int
	Name string
}

// Installation is a single installation
type Installation struct {
	ConnectID            string
	ID                   int
	Name                 string
	Address              string
	Postcode             string
	City                 string
	UsageFunction        string
	ExternalReference    string
	StatusCategory       string
	ServiceOrganisation  ServiceOrganisation
	CustomerOrganisation CustomerOrganisation
}
