package models

// Inventoryreceivingvoucher [...]
type Inventoryreceivingvoucher struct {
	ID         int      `gorm:"primaryKey;column:ID"`
	Createat   int      `gorm:"column:CreateAt"`
	Status     int      `gorm:"column:Status"`
	Totalprice float32  `gorm:"column:TotalPrice"`
	IDprovider int      `gorm:"column:IDProvider"`
	Provider   Provider `gorm:"joinForeignKey:IDProvider;foreignKey:ID;references:IDprovider"`
	IDcreater  int      `gorm:"column:IDCreater"`
	IDimporter int      `gorm:"column:IDImporter"`
}

// TableName get sql table name.获取数据库表名
func (m *Inventoryreceivingvoucher) TableName() string {
	return "InventoryReceivingVoucher"
}
