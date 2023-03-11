package models

// Promotion [...]
type Promotion struct {
	ID        int     `gorm:"primaryKey;column:ID"`
	Createat  int     `gorm:"column:CreateAt"`
	Startat   int     `gorm:"column:StartAt"`
	Endat     int     `gorm:"column:EndAt"`
	Price     float32 `gorm:"column:Price"`
	Productid int     `gorm:"column:ProductID"`
	Product   Product `gorm:"joinForeignKey:ProductID;foreignKey:ID;references:Productid"`
}

// TableName get sql table name.获取数据库表名
func (m *Promotion) TableName() string {
	return "Promotion"
}
