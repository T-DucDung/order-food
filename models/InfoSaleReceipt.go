package models

/******sql******
CREATE TABLE `InfoSaleReceipt` (
  `IDSaleReceipt` int NOT NULL,
  `IDProduct` int NOT NULL,
  `Quantity` float NOT NULL,
  `Price` float NOT NULL,
  `PriceSale` float DEFAULT NULL,
  PRIMARY KEY (`IDSaleReceipt`,`IDProduct`),
  KEY `InfoCart_FK_1` (`IDProduct`) USING BTREE,
  CONSTRAINT `InfoSaleReceipt_FK` FOREIGN KEY (`IDProduct`) REFERENCES `Product` (`ID`),
  CONSTRAINT `InfoSaleReceipt_FK_1` FOREIGN KEY (`IDSaleReceipt`) REFERENCES `SaleReceipt` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Infosalereceipt [...]
type Infosalereceipt struct {
	IDsalereceipt int         `gorm:"primaryKey;column:IDSaleReceipt" json:"-"`
	Salereceipt   Salereceipt `gorm:"joinForeignKey:IDSaleReceipt;foreignKey:ID;references:IDsalereceipt" json:"saleReceiptList"`
	IDproduct     int         `gorm:"primaryKey;column:IDProduct" json:"-"`
	Product       Product     `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct" json:"productList"`
	Quantity      float32     `gorm:"column:Quantity" json:"quantity"`
	Price         float32     `gorm:"column:Price" json:"price"`
	Pricesale     *float32    `gorm:"column:PriceSale" json:"priceSale"`
}

// TableName get sql table name.获取数据库表名
func (m *Infosalereceipt) TableName() string {
	return "InfoSaleReceipt"
}

// InfosalereceiptColumns get sql column name.获取数据库列名
var InfosalereceiptColumns = struct {
	IDsalereceipt string
	IDproduct     string
	Quantity      string
	Price         string
	Pricesale     string
}{
	IDsalereceipt: "IDSaleReceipt",
	IDproduct:     "IDProduct",
	Quantity:      "Quantity",
	Price:         "Price",
	Pricesale:     "PriceSale",
}
