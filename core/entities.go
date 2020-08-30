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
