package admin

import "context"

type Service interface {
	GetItems(context.Context) ([]string, error)
}
