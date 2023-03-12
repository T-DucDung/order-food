package models

/******sql******
CREATE TABLE `Admin` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `UserName` varchar(20) NOT NULL,
  `Pass` varchar(20) NOT NULL,
  `Phone` int DEFAULT NULL,
  `Address` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `Email` varchar(100) DEFAULT NULL,
  `Sex` tinyint(1) DEFAULT NULL,
  `Status` tinyint(1) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Admin [...]
type Admin struct {
	ID       int     `gorm:"primaryKey;column:ID" json:"-"`
	Name     string  `gorm:"column:Name" json:"name"`
	Username string  `gorm:"column:UserName" json:"userName"`
	Pass     string  `gorm:"column:Pass" json:"pass"`
	Phone    *int    `gorm:"column:Phone" json:"phone"`
	Address  *string `gorm:"column:Address" json:"address"`
	Email    *string `gorm:"column:Email" json:"email"`
	Sex      *bool   `gorm:"column:Sex" json:"sex"`
	Status   bool    `gorm:"column:Status" json:"status"`
}

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "Admin"
}

// AdminColumns get sql column name.获取数据库列名
var AdminColumns = struct {
	ID       string
	Name     string
	Username string
	Pass     string
	Phone    string
	Address  string
	Email    string
	Sex      string
	Status   string
}{
	ID:       "ID",
	Name:     "Name",
	Username: "UserName",
	Pass:     "Pass",
	Phone:    "Phone",
	Address:  "Address",
	Email:    "Email",
	Sex:      "Sex",
	Status:   "Status",
}
