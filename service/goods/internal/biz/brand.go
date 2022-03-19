package biz

import (
	"context"
	"errors"
	"goods/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type BrandRepo interface {
	Create(context.Context, *domain.Brand) (*domain.Brand, error)
	IsBrandByID(context.Context, int32) (*domain.Brand, error)
	GetBradByName(context.Context, string) (*domain.Brand, error)
	Update(context.Context, *domain.Brand) error
}

type BrandUsecase struct {
	repo BrandRepo
	log  *log.Helper
}

func NewBrandUsecase(repo BrandRepo, logger log.Logger) *BrandUsecase {
	return &BrandUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BrandUsecase) CreateBrand(ctx context.Context, b *domain.Brand) (*domain.Brand, error) {
	_, err := uc.repo.GetBradByName(ctx, b.Name)
	if err != nil {
		return uc.repo.Create(ctx, b)
	} else {
		return nil, errors.New("当前品牌已经存在")
	}
}

func (uc *BrandUsecase) UpdateBrand(ctx context.Context, b *domain.Brand) error {
	err := uc.repo.Update(ctx, b)
	if err != nil {
		return err
	}
	return nil
}
