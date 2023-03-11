package models

// Address [...]
type Address struct {
	ID         int    `gorm:"primaryKey;column:ID"`
	Name       string `gorm:"column:Name"`
	Phone      string `gorm:"column:Phone"`
	Address    string `gorm:"column:Address"`
	IDcustomer int    `gorm:"column:IDCustomer"`
	User       User   `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer"`
}

// TableName get sql table name.获取数据库表名
func (m *Address) TableName() string {
	return "Address"
}
