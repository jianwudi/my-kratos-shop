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
	s   *biz.SpecificationUsecase
	log *log.Helper
}

func NewGoodsService(cac *biz.CategoryUsecase, gt *biz.GoodsTypeUsecase, s *biz.SpecificationUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{cac: cac, gt: gt, s: s, log: log.NewHelper(logger)}
}
