package read

import "context"

func (WritePostgres) Read(context.Context) ([]string, error) {
	return nil, nil
}
