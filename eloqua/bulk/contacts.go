package bulk

import (
	"context"
)

type ContactsService service

func (s *ContactsService) CreateExport(ctx context.Context, export *Export) (*Export, error) {
	req, err := s.client.NewRequest("POST", "/contacts/exports", export)
	if err != nil {
		return nil, err
	}

	r := &Export{}
	_, err = s.client.Do(ctx, req, r)
	if err != nil {
		return nil, err
	}

	return r, nil

}
