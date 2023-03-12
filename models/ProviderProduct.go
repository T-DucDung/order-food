package models

/******sql******
CREATE TABLE `ProviderProduct` (
  `IDProduct` int NOT NULL,
  `IDProvider` int NOT NULL,
  `Price` float DEFAULT NULL,
  PRIMARY KEY (`IDProduct`,`IDProvider`),
  KEY `ProviderProduct_FK_1` (`IDProvider`),
  CONSTRAINT `ProviderProduct_FK` FOREIGN KEY (`IDProduct`) REFERENCES `Product` (`ID`),
  CONSTRAINT `ProviderProduct_FK_1` FOREIGN KEY (`IDProvider`) REFERENCES `Provider` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Providerproduct [...]
type Providerproduct struct {
	IDproduct  int      `gorm:"primaryKey;column:IDProduct" json:"-"`
	Product    Product  `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct" json:"productList"`
	IDprovider int      `gorm:"primaryKey;column:IDProvider" json:"-"`
	Provider   Provider `gorm:"joinForeignKey:IDProvider;foreignKey:ID;references:IDprovider" json:"providerList"`
	Price      *float32 `gorm:"column:Price" json:"price"`
}

// TableName get sql table name.获取数据库表名
func (m *Providerproduct) TableName() string {
	return "ProviderProduct"
}

// ProviderproductColumns get sql column name.获取数据库列名
var ProviderproductColumns = struct {
	IDproduct  string
	IDprovider string
	Price      string
}{
	IDproduct:  "IDProduct",
	IDprovider: "IDProvider",
	Price:      "Price",
}
