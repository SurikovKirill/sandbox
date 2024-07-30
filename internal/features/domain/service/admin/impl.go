package admin

import "context"

type AdminService struct {
}

func New() (*AdminService, error) {
	return nil, nil
}

func (AdminService) GetItems(ctx context.Context) ([]string, error) {
	return nil, nil
}
