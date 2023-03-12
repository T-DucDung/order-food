package models

/******sql******
CREATE TABLE `SaleReceipt` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `CreateAt` int NOT NULL,
  `DeliveryDate` int DEFAULT NULL,
  `IsPayment` tinyint(1) NOT NULL,
  `Note` varchar(100) DEFAULT NULL,
  `TypePayment` varchar(20) DEFAULT NULL,
  `Status` int NOT NULL,
  `Total` float NOT NULL,
  `IDCustomer` int NOT NULL,
  `IDDeliver` int NOT NULL,
  `RankDiscount` float NOT NULL,
  `IDReceiver` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `SaleReceipt_FK` (`IDCustomer`),
  KEY `SaleReceipt_FK_1` (`IDReceiver`),
  CONSTRAINT `SaleReceipt_FK` FOREIGN KEY (`IDCustomer`) REFERENCES `User` (`ID`),
  CONSTRAINT `SaleReceipt_FK_1` FOREIGN KEY (`IDReceiver`) REFERENCES `Address` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Salereceipt [...]
type Salereceipt struct {
	ID           int     `gorm:"primaryKey;column:ID" json:"-"`
	Createat     int     `gorm:"column:CreateAt" json:"createAt"`
	Deliverydate *int    `gorm:"column:DeliveryDate" json:"deliveryDate"`
	Ispayment    bool    `gorm:"column:IsPayment" json:"isPayment"`
	Note         *string `gorm:"column:Note" json:"note"`
	Typepayment  *string `gorm:"column:TypePayment" json:"typePayment"`
	Status       int     `gorm:"column:Status" json:"status"`
	Total        float32 `gorm:"column:Total" json:"total"`
	IDcustomer   int     `gorm:"column:IDCustomer" json:"idCustomer"`
	User         User    `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer" json:"userList"`
	IDdeliver    int     `gorm:"column:IDDeliver" json:"idDeliver"`
	Rankdiscount float32 `gorm:"column:RankDiscount" json:"rankDiscount"`
	IDreceiver   int     `gorm:"column:IDReceiver" json:"idReceiver"`
	Address      Address `gorm:"joinForeignKey:IDReceiver;foreignKey:ID;references:IDreceiver" json:"addressList"`
}

// TableName get sql table name.获取数据库表名
func (m *Salereceipt) TableName() string {
	return "SaleReceipt"
}

// SalereceiptColumns get sql column name.获取数据库列名
var SalereceiptColumns = struct {
	ID           string
	Createat     string
	Deliverydate string
	Ispayment    string
	Note         string
	Typepayment  string
	Status       string
	Total        string
	IDcustomer   string
	IDdeliver    string
	Rankdiscount string
	IDreceiver   string
}{
	ID:           "ID",
	Createat:     "CreateAt",
	Deliverydate: "DeliveryDate",
	Ispayment:    "IsPayment",
	Note:         "Note",
	Typepayment:  "TypePayment",
	Status:       "Status",
	Total:        "Total",
	IDcustomer:   "IDCustomer",
	IDdeliver:    "IDDeliver",
	Rankdiscount: "RankDiscount",
	IDreceiver:   "IDReceiver",
}
