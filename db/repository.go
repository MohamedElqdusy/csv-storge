package db

import (
	"context"

	"csv-storage/models"
)

type Repository interface {
	CreatePromotion(ctx context.Context, promotion models.Promotion) error
	FindPromotionById(ctx context.Context, promotionId string) (models.Promotion, error)
	Close()
}

var repositoryImpl Repository

func SetRepository(repository Repository) {
	repositoryImpl = repository
}

func CreatePromotion(ctx context.Context, promotion models.Promotion) error {
	return repositoryImpl.CreatePromotion(ctx, promotion)
}

func FindPromotionById(ctx context.Context, promotionId string) (models.Promotion, error) {
	return repositoryImpl.FindPromotionById(ctx, promotionId)

}

func Close() {
	repositoryImpl.Close()
}
