package data

import (
	"context"
	"goods/internal/biz"
	"goods/internal/domain"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

func (esGoodsRepo) GetIndexName() string {
	return "goods"
}
func (esGoodsRepo) GetMapping() string {
	goodsMapping := `
    {
    "mappings": {
        "properties": {
            "id": {
                "type": "integer"
            }
        }
    }
}`
	return goodsMapping
}

// GetMapping 设计商品的 mapping 结构
/* func (esGoodsRepo) GetMapping() string {
	goodsMapping := `
    {
    "mappings": {
        "properties": {
            "id": {
                "type": "integer"
            },
             "brands_id": {
                "type": "integer"
            },
            "category_id": {
                "type": "integer"
            },
            "type_id": {
                "type": "integer"
            },
            "click_num": {
                "type": "integer"
            },
            "fav_num": {
                "type": "integer"
            },
            "is_hot": {
                "type": "boolean"
            },
            "is_new": {
                "type": "boolean"
            },
            "market_price": {
                "type": "integer"
            },
            "name": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "brand_name": {
                "type": "keyword",
                "index": false,
                "dec_values": false,
            },
            "category_name": {
                "type": "keyword",
                "index": false,
                "dec_values": false,
            },
            "type_name": {
                "type": "keyword",
                "index": false,
                "dec_values": false,
            },
            "goods_brief": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "on_sale": {
                "type": "boolean"
            },
            "ship_free": {
                "type": "boolean"
            },
            "shop_price": {
                "type": "integer"
            },
            "sold_num": {
                "type": "integer"
            },
            "sku": {
                "type": "nested",
                "sku_id": {
                    "type": "integer",
                },
                "sku_name": {
                    "type": "text",
                    "analyzer": "ik_max_word"
                },
                "sku_price": {
                    "type": "integer",
                },
            }
        }
    }
}`
	return goodsMapping
} */

type esGoodsRepo struct {
	data *Data
	log  *log.Helper
}

func NewEsGoodsRepo(data *Data, logger log.Logger) biz.EsGoodsRepo {
	return &esGoodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p esGoodsRepo) InsertEsGoods(ctx context.Context, esModel domain.ESGoods) error {
	// 新建 mapping 和 index
	exists, err := p.data.EsClient.IndexExists(p.GetIndexName()).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err = p.data.EsClient.CreateIndex(p.GetIndexName()).BodyString(p.GetMapping()).Do(ctx)
		if err != nil {
			return err
		}
	}

	_, err = p.data.EsClient.Index().Index(p.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(esModel.ID))).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
