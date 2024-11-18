package admin

import "context"

type Transport interface {
	Create(ctx context.Context)
	Read()
	Update()
	Delete()
}
