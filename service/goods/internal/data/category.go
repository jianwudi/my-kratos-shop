package data

import (
	"context"
	"errors"
	"goods/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// Category 商品分类表
type Category struct {
	ID               int32          `gorm:"primarykey;type:int" json:"id"`
	Name             string         `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32          `json:"parent_id"`
	ParentCategory   *Category      `json:"-"`
	SubCategory      []*Category    `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32          `gorm:"column:level;default:1;not null;type:int;comment:分类的级别" json:"level"`
	IsTab            bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort             int32          `gorm:"comment:分类排序;default:99;not null;type:int" json:"sort"`
	CreatedAt        time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type CategoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &CategoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CategoryRepo) AddCategory(ctx context.Context, req *biz.CategoryInfo) (*biz.CategoryInfo, error) {
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab
	cMap["sort"] = req.Sort
	cMap["add_time"] = time.Now()
	cMap["update_time"] = time.Now()

	// 去查询父类目是否存在
	if req.Level != 1 {
		var categories Category
		if res := r.data.db.First(&categories, req.ParentCategory); res.RowsAffected == 0 {
			return nil, errors.New("商品分类不存在")
		}
		cMap["parent_category_id"] = req.ParentCategory
	}

	result := r.data.db.Model(&Category{}).Create(&cMap)
	if result.Error != nil {
		return nil, result.Error
	}
	var value int32
	value, ok := cMap["parent_category_id"].(int32)
	if !ok {
		value = 0
	}
	res := &biz.CategoryInfo{
		Name:           cMap["name"].(string),
		ParentCategory: value,
		Level:          cMap["level"].(int32),
		IsTab:          cMap["is_tab"].(bool),
		Sort:           cMap["sort"].(int32),
	}
	return res, nil
}

func (r *CategoryRepo) GetCategoryByID(ctx context.Context, id int32) (*biz.CategoryInfo, error) {
	var categories Category
	if res := r.data.db.First(&categories, id); res.RowsAffected == 0 {
		return nil, errors.New("商品分类不存在")
	}

	info := &biz.CategoryInfo{
		ID:             categories.ID,
		Name:           categories.Name,
		ParentCategory: categories.ParentCategoryID,
		Level:          categories.Level,
		IsTab:          categories.IsTab,
		Sort:           categories.Sort,
	}
	return info, nil
}
