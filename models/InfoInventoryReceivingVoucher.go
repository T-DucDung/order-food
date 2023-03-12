package models

/******sql******
CREATE TABLE `InfoInventoryReceivingVoucher` (
  `IDProduct` int NOT NULL,
  `IDVoucher` int NOT NULL,
  `Total` int NOT NULL,
  `Price` float NOT NULL,
  PRIMARY KEY (`IDProduct`,`IDVoucher`),
  KEY `InfoInventoryReceivingVoucher_FK` (`IDVoucher`),
  CONSTRAINT `InfoInventoryReceivingVoucher_FK` FOREIGN KEY (`IDVoucher`) REFERENCES `InventoryReceivingVoucher` (`ID`),
  CONSTRAINT `InfoInventoryReceivingVoucher_FK_1` FOREIGN KEY (`IDProduct`) REFERENCES `Product` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Infoinventoryreceivingvoucher [...]
type Infoinventoryreceivingvoucher struct {
	IDproduct                 int                       `gorm:"primaryKey;column:IDProduct" json:"-"`
	Product                   Product                   `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct" json:"productList"`
	IDvoucher                 int                       `gorm:"primaryKey;column:IDVoucher" json:"-"`
	Inventoryreceivingvoucher Inventoryreceivingvoucher `gorm:"joinForeignKey:IDVoucher;foreignKey:ID;references:IDvoucher" json:"inventoryReceivingVoucherList"`
	Total                     int                       `gorm:"column:Total" json:"total"`
	Price                     float32                   `gorm:"column:Price" json:"price"`
}

// TableName get sql table name.获取数据库表名
func (m *Infoinventoryreceivingvoucher) TableName() string {
	return "InfoInventoryReceivingVoucher"
}

// InfoinventoryreceivingvoucherColumns get sql column name.获取数据库列名
var InfoinventoryreceivingvoucherColumns = struct {
	IDproduct string
	IDvoucher string
	Total     string
	Price     string
}{
	IDproduct: "IDProduct",
	IDvoucher: "IDVoucher",
	Total:     "Total",
	Price:     "Price",
}
