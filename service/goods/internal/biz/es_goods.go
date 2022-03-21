package biz

import (
	"context"
	"goods/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type EsGoodsRepo interface {
	InsertEsGoods(context.Context, domain.ESGoods) error // 存储 es 的方法
}

type EsGoodsUsecase struct {
	repo         GoodsRepo
	esRepo       EsGoodsRepo
	categoryRepo CategoryRepo
	log          *log.Helper
}

func NewEsGoodsUsecase(repo GoodsRepo, es EsGoodsRepo, cRepo CategoryRepo, logger log.Logger) *EsGoodsUsecase {
	return &EsGoodsUsecase{
		repo:         repo,  // 商品的repo
		esRepo:       es,    // es 商品的repo
		categoryRepo: cRepo, // 分类的 repo
		log:          log.NewHelper(logger),
	}
}
