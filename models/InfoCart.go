package models

/******sql******
CREATE TABLE `InfoCart` (
  `IDCart` int NOT NULL,
  `IDProduct` int NOT NULL,
  `Quantity` float NOT NULL,
  `Price` float NOT NULL,
  `PriceSale` float DEFAULT NULL,
  PRIMARY KEY (`IDCart`,`IDProduct`),
  KEY `InfoCart_FK_1` (`IDProduct`),
  CONSTRAINT `InfoCart_FK` FOREIGN KEY (`IDCart`) REFERENCES `Cart` (`ID`),
  CONSTRAINT `InfoCart_FK_1` FOREIGN KEY (`IDProduct`) REFERENCES `Product` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Infocart [...]
type Infocart struct {
	IDcart    int      `gorm:"primaryKey;column:IDCart" json:"-"`
	Cart      Cart     `gorm:"joinForeignKey:IDCart;foreignKey:ID;references:IDcart" json:"cartList"`
	IDproduct int      `gorm:"primaryKey;column:IDProduct" json:"-"`
	Product   Product  `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct" json:"productList"`
	Quantity  float32  `gorm:"column:Quantity" json:"quantity"`
	Price     float32  `gorm:"column:Price" json:"price"`
	Pricesale *float32 `gorm:"column:PriceSale" json:"priceSale"`
}

// TableName get sql table name.获取数据库表名
func (m *Infocart) TableName() string {
	return "InfoCart"
}

// InfocartColumns get sql column name.获取数据库列名
var InfocartColumns = struct {
	IDcart    string
	IDproduct string
	Quantity  string
	Price     string
	Pricesale string
}{
	IDcart:    "IDCart",
	IDproduct: "IDProduct",
	Quantity:  "Quantity",
	Price:     "Price",
	Pricesale: "PriceSale",
}
