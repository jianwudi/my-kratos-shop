package service

import (
	v1 "goods/api/goods/v1"
	"goods/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGoodsService)

type GoodsService struct {
	v1.UnimplementedGoodsServer
	cac *biz.CategoryUsecase
	gt  *biz.GoodsTypeUsecase
	bc  *biz.BrandUsecase
	s   *biz.SpecificationUsecase
	ga  *biz.GoodsAttrUsecase
	g   *biz.GoodsUsecase
	log *log.Helper
}

func NewGoodsService(bc *biz.BrandUsecase, cac *biz.CategoryUsecase, gt *biz.GoodsTypeUsecase,
	s *biz.SpecificationUsecase, ga *biz.GoodsAttrUsecase, gc *biz.GoodsUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		bc:  bc,
		cac: cac,
		gt:  gt,
		s:   s,
		ga:  ga,
		g:   gc,
		log: log.NewHelper(logger),
	}
}
