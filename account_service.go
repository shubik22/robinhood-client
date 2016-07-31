package client

import (
	"net/http"

	"github.com/shubik22/go-robinhood/lib/models"
)

type AccountService service

func (s *AccountService) ListAccounts() (*models.AccountsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "accounts/", nil)
	if err != nil {
		return nil, nil, err
	}

	a := &models.AccountsResponse{}
	resp, err := s.client.Do(req, a)

	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
