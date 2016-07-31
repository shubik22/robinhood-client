package robinhood

import (
	"net/http"
)

type PositionService service

func (s *PositionService) ListPositions() (*PositionsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "positions/", nil)
	if err != nil {
		return nil, nil, err
	}

	p := &PositionsResponse{}
	resp, err := s.client.Do(req, p)

	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}
