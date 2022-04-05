package ports

import (
	"fmt"

	"github.com/txya900619/url-shortener/internal/kgs/schema"
	"github.com/txya900619/url-shortener/pkg/queue"

	"context"

	kgs_grpc "github.com/txya900619/url-shortener/pkg/genproto/kgs"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type KeyServiceServer struct {
	kgs_grpc.UnimplementedKeyServiceServer
	cacheQueue *queue.StringQueue
	db         *gorm.DB
}

func NewKeyServiceServer(db *gorm.DB) (*KeyServiceServer, error) {
	cacheQueue := queue.NewStringQueue(100)
	for i := 0; i < 100; i++ {
		key, err := generateKey(db)
		if err != nil {
			return nil, err
		}

		err = cacheQueue.Insert(key)
		if err != nil {
			return nil, err
		}
	}

	return &KeyServiceServer{db: db, cacheQueue: cacheQueue}, nil
}

func (s *KeyServiceServer) insertKey() error {
	key, err := generateKey(s.db)
	if err != nil {
		fmt.Printf("insertKey error: %v", err)
		return s.insertKey()
	}

	s.cacheQueue.Insert(key)
	return nil
}

func (s *KeyServiceServer) GenerateKey(ctx context.Context, in *emptypb.Empty) (*kgs_grpc.GenerateKeyResponse, error) {
	key, err := s.cacheQueue.Remove()
	if err != nil {
		key, err = generateKey(s.db)
		if err != nil {
			return nil, err
		}
	}

	go s.insertKey()

	return &kgs_grpc.GenerateKeyResponse{Key: key}, nil
}

func (s *KeyServiceServer) DeleteKeys(ctx context.Context, in *kgs_grpc.DeleteKeyRequest) (*emptypb.Empty, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, key := range in.Keys {
			if err := s.db.Delete(&schema.UsedKey{Key: key}).Error; err != nil {
				tx.Rollback()
				return err
			}

			if err := s.db.Create(&schema.UnusedKey{Key: key}).Error; err != nil {
				tx.Rollback()
				return err
			}

		}

		return tx.Commit().Error
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func generateKey(db *gorm.DB) (string, error) {
	var unusedKey schema.UnusedKey

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Order("random()").First(&unusedKey).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Delete(&unusedKey).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&schema.UsedKey{Key: unusedKey.Key}).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	})

	if err != nil {
		return "", err
	}

	return unusedKey.Key, nil
}
