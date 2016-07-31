package robinhood

import (
	"net/http"
)

type AccountService service

func (s *AccountService) ListAccounts() (*AccountsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "accounts/", nil)
	if err != nil {
		return nil, nil, err
	}

	a := &AccountsResponse{}
	resp, err := s.client.Do(req, a)

	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
