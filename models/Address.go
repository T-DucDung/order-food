package models

/******sql******
CREATE TABLE `Address` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  `Phone` varchar(20) NOT NULL,
  `Address` varchar(100) NOT NULL,
  `IDCustomer` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `Address_FK` (`IDCustomer`),
  CONSTRAINT `Address_FK` FOREIGN KEY (`IDCustomer`) REFERENCES `User` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Address [...]
type Address struct {
	ID         int    `gorm:"primaryKey;column:ID" json:"-"`
	Name       string `gorm:"column:Name" json:"name"`
	Phone      string `gorm:"column:Phone" json:"phone"`
	Address    string `gorm:"column:Address" json:"address"`
	IDcustomer int    `gorm:"column:IDCustomer" json:"idCustomer"`
	User       User   `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer" json:"userList"`
}

// TableName get sql table name.获取数据库表名
func (m *Address) TableName() string {
	return "Address"
}

// AddressColumns get sql column name.获取数据库列名
var AddressColumns = struct {
	ID         string
	Name       string
	Phone      string
	Address    string
	IDcustomer string
}{
	ID:         "ID",
	Name:       "Name",
	Phone:      "Phone",
	Address:    "Address",
	IDcustomer: "IDCustomer",
}
