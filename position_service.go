package client

import (
	"net/http"

	"github.com/shubik22/go-robinhood/lib/models"
)

type PositionService service

func (s *PositionService) ListPositions() (*models.PositionsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "positions/", nil)
	if err != nil {
		return nil, nil, err
	}

	p := &models.PositionsResponse{}
	resp, err := s.client.Do(req, p)

	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}
