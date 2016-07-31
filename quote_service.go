package client

import (
	"net/http"

	"github.com/shubik22/go-robinhood/lib/models"
)

type QuoteService service

func (s *QuoteService) GetQuote(p *models.Position) (*models.Quote, *http.Response, error) {
	i, _, err := s.getInstrument(p)
	if err != nil {
		return nil, nil, err
	}

	return s.getQuote(i)
}

func (s *QuoteService) getInstrument(p *models.Position) (*models.Instrument, *http.Response, error) {
	req, err := s.client.NewRequestWithFullUrl("GET", p.Instrument, nil)
	if err != nil {
		return nil, nil, err
	}

	i := &models.Instrument{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}

func (s *QuoteService) getQuote(i *models.Instrument) (*models.Quote, *http.Response, error) {
	req, err := s.client.NewRequestWithFullUrl("GET", i.Quote, nil)
	if err != nil {
		return nil, nil, err
	}

	q := &models.Quote{}
	resp, err := s.client.Do(req, q)

	if err != nil {
		return nil, resp, err
	}

	return q, resp, err
}
