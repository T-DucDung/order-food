package models

// Cart [...]
type Cart struct {
	ID         int     `gorm:"primaryKey;column:ID"`
	IDcustomer int     `gorm:"column:IDCustomer"`
	User       User    `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer"`
	Total      float32 `gorm:"column:Total"`
}

// TableName get sql table name.获取数据库表名
func (m *Cart) TableName() string {
	return "Cart"
}
