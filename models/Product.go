package models

// Product [...]
type Product struct {
	ID          int         `gorm:"primaryKey;column:ID"`
	Name        string      `gorm:"column:Name"`
	Img         string      `gorm:"column:Img"`
	Price       float32     `gorm:"column:Price"`
	Unit        string      `gorm:"column:unit"`
	Remain      float32     `gorm:"column:Remain"`
	Sold        float32     `gorm:"column:Sold"`
	Description string      `gorm:"column:Description"`
	Rateavg     int         `gorm:"column:RateAVG"`
	IDtype      int         `gorm:"column:IDType"`
	Producttype Producttype `gorm:"joinForeignKey:IDType;foreignKey:ID;references:IDtype"`
}

// TableName get sql table name.获取数据库表名
func (m *Product) TableName() string {
	return "Product"
}
