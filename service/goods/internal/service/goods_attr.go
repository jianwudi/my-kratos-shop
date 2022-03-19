package service

import (
	"context"

	v1 "goods/api/goods/v1"
	"goods/internal/domain"
)

// CreateAttrGroup 创建属性组
func (g *GoodsService) CreateAttrGroup(ctx context.Context, r *v1.AttrGroupRequest) (*v1.AttrGroupResponse, error) {
	result, err := g.ga.CreateAttrGroup(ctx, &domain.AttrGroup{
		TypeID: r.TypeId,
		Title:  r.Title,
		Desc:   r.Desc,
		Status: r.Status,
		Sort:   r.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &v1.AttrGroupResponse{
		Id:     result.ID,
		TypeId: result.TypeID,
		Title:  result.Title,
		Desc:   result.Desc,
		Status: result.Status,
		Sort:   result.Sort,
	}, nil
}
