package read

import "context"

type WriteRepository interface {
	Read(context.Context) ([]string, error)
}
