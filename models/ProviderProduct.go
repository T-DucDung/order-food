package models

// Providerproduct [...]
type Providerproduct struct {
	IDproduct  int      `gorm:"primaryKey;column:IDProduct"`
	Product    Product  `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct"`
	IDprovider int      `gorm:"primaryKey;column:IDProvider"`
	Provider   Provider `gorm:"joinForeignKey:IDProvider;foreignKey:ID;references:IDprovider"`
	Price      float32  `gorm:"column:Price"`
}

// TableName get sql table name.获取数据库表名
func (m *Providerproduct) TableName() string {
	return "ProviderProduct"
}
