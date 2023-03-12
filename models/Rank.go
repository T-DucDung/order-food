package models

/******sql******
CREATE TABLE `Rank` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  `TotalSpend` float NOT NULL,
  `Discount` float NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Rank [...]
type Rank struct {
	ID         int     `gorm:"primaryKey;column:ID" json:"-"`
	Name       string  `gorm:"column:Name" json:"name"`
	Totalspend float32 `gorm:"column:TotalSpend" json:"totalSpend"`
	Discount   float32 `gorm:"column:Discount" json:"discount"`
}

// TableName get sql table name.获取数据库表名
func (m *Rank) TableName() string {
	return "Rank"
}

// RankColumns get sql column name.获取数据库列名
var RankColumns = struct {
	ID         string
	Name       string
	Totalspend string
	Discount   string
}{
	ID:         "ID",
	Name:       "Name",
	Totalspend: "TotalSpend",
	Discount:   "Discount",
}
