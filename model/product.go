package model

import "gorm.io/gorm"

type Product struct {
  gorm.Model
  PublishModel PublishModel `gorm:"embedded"`
  SkuId string `json:"skuId" gorm:"size:11;not null;uniqueIndex"`
  SkuName string `json:"skuName" gorm:"type:varchar(255);not null"`
  Price float64 `json:"price" gorm:"not null"`
  PromotionPrice float64 `json:"promotionPrice" gorm:"default:-1;comment:'促销价格'"`
  Stock int `json:"stock" gorm:"type:SMALLINT UNSIGNED not NULL;comment:'库存量'"`
  Status string `json:"status" gorm:"type:enum('published', 'pending', 'deleted');default:'pending';comment:'产品发布状态'"`
}

func ProductIsExistedBySkuId(db *gorm.DB, skuId string) (bool, error) {
  var product Product
  err := db.Select("id").Where("sku_id = ?", skuId).First(&product).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return false, err
  }

  if product.ID > 0 {
    return true, nil
  }

  return false, nil
}

func ProductAddNew(db *gorm.DB, product Product) error {
  err := db.Create(&product).Error
  if err != nil {
    return err
  }

  return nil
}
