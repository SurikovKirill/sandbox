package write

import "context"

type WriteRepository interface {
	Create(context.Context, string) error
	Update(context.Context, string) error
	Delete(context.Context, string) error
}
