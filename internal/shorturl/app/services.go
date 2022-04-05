package app

import "context"

type KeyService interface {
	GenerateKey(ctx context.Context) (string, error)
}
