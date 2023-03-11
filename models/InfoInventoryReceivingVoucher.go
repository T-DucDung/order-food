package models

// Infoinventoryreceivingvoucher [...]
type Infoinventoryreceivingvoucher struct {
	IDproduct                 int                       `gorm:"primaryKey;column:IDProduct"`
	Product                   Product                   `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct"`
	IDvoucher                 int                       `gorm:"primaryKey;column:IDVoucher"`
	Inventoryreceivingvoucher Inventoryreceivingvoucher `gorm:"joinForeignKey:IDVoucher;foreignKey:ID;references:IDvoucher"`
	Total                     int                       `gorm:"column:Total"`
	Price                     float32                   `gorm:"column:Price"`
}

// TableName get sql table name.获取数据库表名
func (m *Infoinventoryreceivingvoucher) TableName() string {
	return "InfoInventoryReceivingVoucher"
}
