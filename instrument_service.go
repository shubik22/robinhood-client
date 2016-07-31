package robinhood

import (
	"fmt"
	"net/http"
)

type InstrumentService service

func (s *InstrumentService) GetInstrumentFromSymbol(symbol string) (*Instrument, *http.Response, error) {
	path := fmt.Sprintf("instruments/?query=%v", symbol)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	ir := &InstrumentsResponse{}
	resp, err := s.client.Do(req, ir)

	if err != nil {
		return nil, resp, err
	}

	i := ir.Results[0]

	return &i, resp, err
}

func (s *InstrumentService) GetInstrumentFromPosition(p *Position) (*Instrument, *http.Response, error) {
	req, err := s.client.NewRequestWithFullUrl("GET", p.Instrument, nil)
	if err != nil {
		return nil, nil, err
	}

	i := &Instrument{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}
