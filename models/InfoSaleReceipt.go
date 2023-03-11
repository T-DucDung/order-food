package models

// Infosalereceipt [...]
type Infosalereceipt struct {
	IDsalereceipt int         `gorm:"primaryKey;column:IDSaleReceipt"`
	Salereceipt   Salereceipt `gorm:"joinForeignKey:IDSaleReceipt;foreignKey:ID;references:IDsalereceipt"`
	IDproduct     int         `gorm:"primaryKey;column:IDProduct"`
	Product       Product     `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct"`
	Quantity      float32     `gorm:"column:Quantity"`
	Price         float32     `gorm:"column:Price"`
	Pricesale     float32     `gorm:"column:PriceSale"`
}

// TableName get sql table name.获取数据库表名
func (m *Infosalereceipt) TableName() string {
	return "InfoSaleReceipt"
}
