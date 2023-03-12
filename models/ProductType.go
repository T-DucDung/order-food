package models

/******sql******
CREATE TABLE `ProductType` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Producttype [...]
type Producttype struct {
	ID   int    `gorm:"primaryKey;column:ID" json:"-"`
	Name string `gorm:"column:Name" json:"name"`
}

// TableName get sql table name.获取数据库表名
func (m *Producttype) TableName() string {
	return "ProductType"
}

// ProducttypeColumns get sql column name.获取数据库列名
var ProducttypeColumns = struct {
	ID   string
	Name string
}{
	ID:   "ID",
	Name: "Name",
}
