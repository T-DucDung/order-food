package models

// User [...]
type User struct {
	ID       int    `gorm:"primaryKey;column:ID"`
	Name     string `gorm:"column:Name"`
	Username string `gorm:"column:Username"`
	Pass     string `gorm:"column:Pass"`
	Phone    string `gorm:"column:Phone"`
	Address  string `gorm:"column:Address"`
	Email    string `gorm:"column:Email"`
	Sex      bool   `gorm:"column:Sex"`
	Createat int    `gorm:"column:CreateAt"`
	Status   bool   `gorm:"column:Status"`
	IDrank   int    `gorm:"column:IDRank"`
	Rank     Rank   `gorm:"joinForeignKey:IDRank;foreignKey:ID;references:IDrank"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "User"
}
