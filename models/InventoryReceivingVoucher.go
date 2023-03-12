package models

/******sql******
CREATE TABLE `InventoryReceivingVoucher` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `CreateAt` int NOT NULL,
  `Status` int NOT NULL,
  `TotalPrice` float NOT NULL,
  `IDProvider` int NOT NULL,
  `IDCreater` int NOT NULL,
  `IDImporter` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `InventoryReceivingVoucher_FK` (`IDProvider`),
  CONSTRAINT `InventoryReceivingVoucher_FK` FOREIGN KEY (`IDProvider`) REFERENCES `Provider` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Inventoryreceivingvoucher [...]
type Inventoryreceivingvoucher struct {
	ID         int      `gorm:"primaryKey;column:ID" json:"-"`
	Createat   int      `gorm:"column:CreateAt" json:"createAt"`
	Status     int      `gorm:"column:Status" json:"status"`
	Totalprice float32  `gorm:"column:TotalPrice" json:"totalPrice"`
	IDprovider int      `gorm:"column:IDProvider" json:"idProvider"`
	Provider   Provider `gorm:"joinForeignKey:IDProvider;foreignKey:ID;references:IDprovider" json:"providerList"`
	IDcreater  int      `gorm:"column:IDCreater" json:"idCreater"`
	IDimporter int      `gorm:"column:IDImporter" json:"idImporter"`
}

// TableName get sql table name.获取数据库表名
func (m *Inventoryreceivingvoucher) TableName() string {
	return "InventoryReceivingVoucher"
}

// InventoryreceivingvoucherColumns get sql column name.获取数据库列名
var InventoryreceivingvoucherColumns = struct {
	ID         string
	Createat   string
	Status     string
	Totalprice string
	IDprovider string
	IDcreater  string
	IDimporter string
}{
	ID:         "ID",
	Createat:   "CreateAt",
	Status:     "Status",
	Totalprice: "TotalPrice",
	IDprovider: "IDProvider",
	IDcreater:  "IDCreater",
	IDimporter: "IDImporter",
}
