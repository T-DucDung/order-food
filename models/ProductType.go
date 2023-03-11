package models

// Producttype [...]
type Producttype struct {
	ID   int    `gorm:"primaryKey;column:ID"`
	Name string `gorm:"column:Name"`
}

// TableName get sql table name.获取数据库表名
func (m *Producttype) TableName() string {
	return "ProductType"
}
