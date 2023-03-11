package models

// Admin [...]
type Admin struct {
	ID       int    `gorm:"primaryKey;column:ID"`
	Name     string `gorm:"column:Name"`
	Username string `gorm:"column:UserName"`
	Pass     string `gorm:"column:Pass"`
	Phone    int    `gorm:"column:Phone"`
	Address  string `gorm:"column:Address"`
	Email    string `gorm:"column:Email"`
	Sex      bool   `gorm:"column:Sex"`
	Status   bool   `gorm:"column:Status"`
}

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "Admin"
}
