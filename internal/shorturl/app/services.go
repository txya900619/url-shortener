package app

import "context"

type KeyService interface {
	GenerateKey(ctx context.Context) (string, error)
	DeleteKeys(ctx context.Context, keys []string) error
}
