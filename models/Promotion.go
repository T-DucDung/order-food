package models

/******sql******
CREATE TABLE `Promotion` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `CreateAt` int NOT NULL,
  `StartAt` int NOT NULL,
  `EndAt` int NOT NULL,
  `Price` float NOT NULL,
  `ProductID` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `SaleInfo_FK` (`ProductID`),
  CONSTRAINT `SaleInfo_FK` FOREIGN KEY (`ProductID`) REFERENCES `Product` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Promotion [...]
type Promotion struct {
	ID        int     `gorm:"primaryKey;column:ID" json:"-"`
	Createat  int     `gorm:"column:CreateAt" json:"createAt"`
	Startat   int     `gorm:"column:StartAt" json:"startAt"`
	Endat     int     `gorm:"column:EndAt" json:"endAt"`
	Price     float32 `gorm:"column:Price" json:"price"`
	Productid int     `gorm:"column:ProductID" json:"productId"`
	Product   Product `gorm:"joinForeignKey:ProductID;foreignKey:ID;references:Productid" json:"productList"`
}

// TableName get sql table name.获取数据库表名
func (m *Promotion) TableName() string {
	return "Promotion"
}

// PromotionColumns get sql column name.获取数据库列名
var PromotionColumns = struct {
	ID        string
	Createat  string
	Startat   string
	Endat     string
	Price     string
	Productid string
}{
	ID:        "ID",
	Createat:  "CreateAt",
	Startat:   "StartAt",
	Endat:     "EndAt",
	Price:     "Price",
	Productid: "ProductID",
}
