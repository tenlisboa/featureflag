package contracts

import "featureflag/internal/pkg/entities"

type Repository interface {
	List() ([]entities.Flag, error)
	Get(key string) (entities.Flag, error)
	Save(flag entities.Flag) error
}
