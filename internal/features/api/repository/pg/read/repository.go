package read

import "context"

type ReadRepository interface {
	Read(context.Context) ([]string, error)
}
