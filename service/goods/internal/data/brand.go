package data

import (
	"context"
	"goods/internal/biz"
	"goods/internal/domain"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// Brand 商品品牌表
type Brand struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:品牌名称" json:"name"`
	Logo      string         `gorm:"type:varchar(200);default:;comment:品牌Logo图片"`
	Desc      string         `gorm:"type:varchar(500);default:;comment:品牌描述"`
	IsTab     bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort      int32          `gorm:"comment:品牌排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (p *Brand) ToDomain() *domain.Brand {
	return &domain.Brand{
		ID:    p.ID,
		Name:  p.Name,
		Logo:  p.Logo,
		Desc:  p.Desc,
		IsTab: p.IsTab,
		Sort:  p.Sort,
	}
}

type BrandRepo struct {
	data *Data
	log  *log.Helper
}

// NewBrandRepo .
func NewBrandRepo(data *Data, logger log.Logger) biz.BrandRepo {
	return &BrandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *BrandRepo) Create(ctx context.Context, b *domain.Brand) (*domain.Brand, error) {
	brand := &Brand{
		Name:  b.Name,
		Logo:  b.Logo,
		Desc:  b.Desc,
		IsTab: b.IsTab,
		Sort:  b.Sort,
	}
	if err := r.data.db.Save(brand).Error; err != nil {
		return nil, errors.InternalServer("SAVE_BRAND_ERROR", err.Error())
	}
	res := &domain.Brand{
		ID:    brand.ID,
		Name:  brand.Name,
		Logo:  brand.Logo,
		Desc:  brand.Desc,
		IsTab: brand.IsTab,
		Sort:  brand.Sort,
	}
	return res, nil
}

func (r *BrandRepo) IsBrandByID(ctx context.Context, id int32) (*domain.Brand, error) {
	var b Brand
	if err := r.data.db.Table("brands").Where("id = ?", id).First(&b).Error; err != nil {
		return nil, errors.InternalServer("BRAND_NOT_FOUND", err.Error())
	}

	return b.ToDomain(), nil
}

func (r *BrandRepo) GetBradByName(ctx context.Context, name string) (*domain.Brand, error) {
	var brand Brand
	result := r.data.db.Where("name=?", name).First(&brand)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("BRAND_NOT_FOUND", "brand not found")
	}
	if result.RowsAffected == 1 {
		return &domain.Brand{
			ID:    brand.ID,
			Name:  brand.Name,
			Logo:  brand.Logo,
			Desc:  brand.Desc,
			IsTab: brand.IsTab,
			Sort:  brand.Sort,
		}, nil
	} else {
		return nil, errors.NotFound("BRAND_NOT_FOUND", "brand not found")
	}
}

func (r *BrandRepo) Update(ctx context.Context, b *domain.Brand) error {
	brands := Brand{}
	if result := r.data.db.Where("id=?", b.ID).First(&brands); result.RowsAffected == 0 {
		return errors.NotFound("BRAND_NOT_FOUND", "brand not found")
	}

	if b.Name != "" {
		brands.Name = b.Name
	}
	if b.Logo != "" {
		brands.Logo = b.Logo
	}
	if b.IsTab {
		brands.IsTab = b.IsTab
	}
	if b.Sort != 0 {
		brands.Sort = b.Sort
	}
	if b.Desc != "" {
		brands.Desc = b.Desc
	}
	if err := r.data.db.Save(&brands).Error; err != nil {
		return errors.InternalServer("UPDATE_BRAND_ERROR", err.Error())
	}
	return nil
}
