package client

import (
	"net/http"
	"net/url"
)

type authResponse struct {
	Token string `json:"token"`
}

type authRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}

type AuthenticationService service

func (s *AuthenticationService) Login() (*http.Response, error) {
	params := url.Values{}
	params.Add("username", s.client.UserName)
	params.Add("password", s.client.Password)

	a := &authResponse{}
	resp, err := s.client.PostForm("api-token-auth/", params, a)

	if err == nil {
		s.client.AuthToken = a.Token
	}

	return resp, err
}
