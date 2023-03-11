package models

// Provider [...]
type Provider struct {
	ID      int    `gorm:"primaryKey;column:ID"`
	Name    string `gorm:"column:Name"`
	Phone   int    `gorm:"column:Phone"`
	Address string `gorm:"column:Address"`
	Email   string `gorm:"column:Email"`
}

// TableName get sql table name.获取数据库表名
func (m *Provider) TableName() string {
	return "Provider"
}
