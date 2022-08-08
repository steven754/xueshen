package model

import (
	"time"
)

// 商品类型表
type GoodsType struct {
	BaseEntity
	Id         int
	Name       string    // 种类名称
	Logo       string    //logo
	Image      string    // 图片
	GoodsSKUId *GoodsSKU `gorm:"foreignkey:GoodsTypeId"` // 商品销量
}

// 商品SKU表
type GoodsSKU struct {
	BaseEntity
	Id          int
	GoodsId     int `gorm:"foreignkey:Id"` // 商品 SPU
	GoodsTypeId int
	Name        string      // 商品名称
	Description string      // 商品简介
	Price       int         // 商品价格
	Unite       string      // 商品单位
	Image       string      // 商品图片
	Stock       int         `gorm:"default:1"` // 商品库存
	Sales       int         `gorm:"default:0"` // 商品销量
	Status      int         `gorm:"default:1"` // 商品状态
	Time        time.Time   // 添加时间
	GoodsImage  *GoodsImage `gorm:"foreignkey:Id"`
}

// 商品SPU表
type Goods struct {
	BaseEntity
	Id     int
	Name   string `gorm:"type:varchar(20)"`  // 商品名称
	Detail string `gorm:"type:varchar(200)"` // 详细描述
}

// 首页分类商品展示表
type IndexTypeGoodsBanner struct {
	BaseEntity
	Id          int
	GoodsTypeId int // 商品类型
	GoodsSKUId  int // 商品sku
	DisplayType int `gorm:"default:1"` // 展示类型0代表文字，1 代表图片
	IndexId     int `gorm:"default:0"` // 展示顺序
}

// 商品图片表
type GoodsImage struct {
	BaseEntity
	Id    int
	Image string // 商品图片
}
