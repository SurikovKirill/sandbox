package read

import "context"

type ReadRepository interface {
	Create(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
}
