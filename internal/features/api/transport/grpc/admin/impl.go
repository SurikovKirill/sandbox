package admin

import (
	"sandbox-service/internal/features/api/service/admin"
)

type AdminHandler struct {
	service admin.Service
}

func New() (*AdminHandler, error) {
	return nil, nil
}
