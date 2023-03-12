package models

/******sql******
CREATE TABLE `User` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  `Username` varchar(20) NOT NULL,
  `Pass` varchar(20) NOT NULL,
  `Phone` varchar(20) DEFAULT NULL,
  `Address` varchar(20) DEFAULT NULL,
  `Email` varchar(20) DEFAULT NULL,
  `Sex` tinyint(1) NOT NULL,
  `CreateAt` int NOT NULL,
  `Status` tinyint(1) NOT NULL,
  `IDRank` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `User_FK` (`IDRank`),
  CONSTRAINT `User_FK` FOREIGN KEY (`IDRank`) REFERENCES `Rank` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// User [...]
type User struct {
	ID       int     `gorm:"primaryKey;column:ID" json:"-"`
	Name     string  `gorm:"column:Name" json:"name"`
	Username string  `gorm:"column:Username" json:"username"`
	Pass     string  `gorm:"column:Pass" json:"pass"`
	Phone    *string `gorm:"column:Phone" json:"phone"`
	Address  *string `gorm:"column:Address" json:"address"`
	Email    *string `gorm:"column:Email" json:"email"`
	Sex      bool    `gorm:"column:Sex" json:"sex"`
	Createat int     `gorm:"column:CreateAt" json:"createAt"`
	Status   bool    `gorm:"column:Status" json:"status"`
	IDrank   int     `gorm:"column:IDRank" json:"idRank"`
	Rank     Rank    `gorm:"joinForeignKey:IDRank;foreignKey:ID;references:IDrank" json:"rankList"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "User"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID       string
	Name     string
	Username string
	Pass     string
	Phone    string
	Address  string
	Email    string
	Sex      string
	Createat string
	Status   string
	IDrank   string
}{
	ID:       "ID",
	Name:     "Name",
	Username: "Username",
	Pass:     "Pass",
	Phone:    "Phone",
	Address:  "Address",
	Email:    "Email",
	Sex:      "Sex",
	Createat: "CreateAt",
	Status:   "Status",
	IDrank:   "IDRank",
}
