package models

// Infocart [...]
type Infocart struct {
	IDcart    int     `gorm:"primaryKey;column:IDCart"`
	Cart      Cart    `gorm:"joinForeignKey:IDCart;foreignKey:ID;references:IDcart"`
	IDproduct int     `gorm:"primaryKey;column:IDProduct"`
	Product   Product `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct"`
	Quantity  float32 `gorm:"column:Quantity"`
	Price     float32 `gorm:"column:Price"`
	Pricesale float32 `gorm:"column:PriceSale"`
}

// TableName get sql table name.获取数据库表名
func (m *Infocart) TableName() string {
	return "InfoCart"
}
