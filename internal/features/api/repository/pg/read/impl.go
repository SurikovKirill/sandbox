package read

import "context"

type ReadPostgres struct {
}

func New() (*ReadPostgres, error) {
	return nil, nil
}

func (ReadPostgres) Create(context.Context) error {
	return nil
}

func (ReadPostgres) Update(context.Context) error {
	return nil
}
func (ReadPostgres) Delete(context.Context) error {
	return nil
}
