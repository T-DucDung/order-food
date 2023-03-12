package models

/******sql******
CREATE TABLE `Product` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  `Img` varchar(20) NOT NULL,
  `Price` float NOT NULL,
  `unit` varchar(20) NOT NULL,
  `Remain` float NOT NULL,
  `Sold` float NOT NULL,
  `Description` varchar(20) DEFAULT NULL,
  `RateAVG` int NOT NULL,
  `IDType` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `Product_FK` (`IDType`),
  CONSTRAINT `Product_FK` FOREIGN KEY (`IDType`) REFERENCES `ProductType` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Product [...]
type Product struct {
	ID          int         `gorm:"primaryKey;column:ID" json:"-"`
	Name        string      `gorm:"column:Name" json:"name"`
	Img         string      `gorm:"column:Img" json:"img"`
	Price       float32     `gorm:"column:Price" json:"price"`
	Unit        string      `gorm:"column:unit" json:"unit"`
	Remain      float32     `gorm:"column:Remain" json:"remain"`
	Sold        float32     `gorm:"column:Sold" json:"sold"`
	Description *string     `gorm:"column:Description" json:"description"`
	Rateavg     int         `gorm:"column:RateAVG" json:"rateAVG"`
	IDtype      int         `gorm:"column:IDType" json:"idType"`
	Producttype Producttype `gorm:"joinForeignKey:IDType;foreignKey:ID;references:IDtype" json:"productTypeList"`
}

// TableName get sql table name.获取数据库表名
func (m *Product) TableName() string {
	return "Product"
}

// ProductColumns get sql column name.获取数据库列名
var ProductColumns = struct {
	ID          string
	Name        string
	Img         string
	Price       string
	Unit        string
	Remain      string
	Sold        string
	Description string
	Rateavg     string
	IDtype      string
}{
	ID:          "ID",
	Name:        "Name",
	Img:         "Img",
	Price:       "Price",
	Unit:        "unit",
	Remain:      "Remain",
	Sold:        "Sold",
	Description: "Description",
	Rateavg:     "RateAVG",
	IDtype:      "IDType",
}
