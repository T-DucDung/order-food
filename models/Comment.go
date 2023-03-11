package models

// Comment [...]
type Comment struct {
	ID         int     `gorm:"primaryKey;column:ID"`
	Content    string  `gorm:"column:Content"`
	Createat   int     `gorm:"column:CreateAt"`
	IDproduct  int     `gorm:"column:IDProduct"`
	Product    Product `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct"`
	IDcustomer int     `gorm:"column:IDCustomer"`
	User       User    `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer"`
	Rate       int     `gorm:"column:Rate"`
}

// TableName get sql table name.获取数据库表名
func (m *Comment) TableName() string {
	return "Comment"
}
