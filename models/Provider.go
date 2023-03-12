package models

/******sql******
CREATE TABLE `Provider` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  `Phone` int DEFAULT NULL,
  `Address` varchar(20) DEFAULT NULL,
  `Email` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Provider [...]
type Provider struct {
	ID      int     `gorm:"primaryKey;column:ID" json:"-"`
	Name    string  `gorm:"column:Name" json:"name"`
	Phone   *int    `gorm:"column:Phone" json:"phone"`
	Address *string `gorm:"column:Address" json:"address"`
	Email   *string `gorm:"column:Email" json:"email"`
}

// TableName get sql table name.获取数据库表名
func (m *Provider) TableName() string {
	return "Provider"
}

// ProviderColumns get sql column name.获取数据库列名
var ProviderColumns = struct {
	ID      string
	Name    string
	Phone   string
	Address string
	Email   string
}{
	ID:      "ID",
	Name:    "Name",
	Phone:   "Phone",
	Address: "Address",
	Email:   "Email",
}
