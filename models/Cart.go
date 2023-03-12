package models

/******sql******
CREATE TABLE `Cart` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `IDCustomer` int NOT NULL,
  `Total` float NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `Cart_FK` (`IDCustomer`),
  CONSTRAINT `Cart_FK` FOREIGN KEY (`IDCustomer`) REFERENCES `User` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Cart [...]
type Cart struct {
	ID         int     `gorm:"primaryKey;column:ID" json:"-"`
	IDcustomer int     `gorm:"column:IDCustomer" json:"idCustomer"`
	User       User    `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer" json:"userList"`
	Total      float32 `gorm:"column:Total" json:"total"`
}

// TableName get sql table name.获取数据库表名
func (m *Cart) TableName() string {
	return "Cart"
}

// CartColumns get sql column name.获取数据库列名
var CartColumns = struct {
	ID         string
	IDcustomer string
	Total      string
}{
	ID:         "ID",
	IDcustomer: "IDCustomer",
	Total:      "Total",
}
