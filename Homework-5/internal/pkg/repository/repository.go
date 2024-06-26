//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import "context"

type PvzRepo interface {
	Add(ctx context.Context, pvz *Pvz) (int64, error)
	GetByID(ctx context.Context, id int64) (*Pvz, error)
	Update(ctx context.Context, pvz *Pvz) error
	DeleteByID(ctx context.Context, id int64) error
}
