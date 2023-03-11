package models

// Rank [...]
type Rank struct {
	ID         int     `gorm:"primaryKey;column:ID"`
	Name       string  `gorm:"column:Name"`
	Totalspend float32 `gorm:"column:TotalSpend"`
	Discount   float32 `gorm:"column:Discount"`
}

// TableName get sql table name.获取数据库表名
func (m *Rank) TableName() string {
	return "Rank"
}
