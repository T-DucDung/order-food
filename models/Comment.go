package models

/******sql******
CREATE TABLE `Comment` (
  `ID` int NOT NULL,
  `Content` varchar(20) NOT NULL,
  `CreateAt` int NOT NULL,
  `IDProduct` int NOT NULL,
  `IDCustomer` int NOT NULL,
  `Rate` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `Comment_FK` (`IDProduct`),
  KEY `Comment_FK_1` (`IDCustomer`),
  CONSTRAINT `Comment_FK` FOREIGN KEY (`IDProduct`) REFERENCES `Product` (`ID`),
  CONSTRAINT `Comment_FK_1` FOREIGN KEY (`IDCustomer`) REFERENCES `User` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Comment [...]
type Comment struct {
	ID         int     `gorm:"primaryKey;column:ID" json:"-"`
	Content    string  `gorm:"column:Content" json:"content"`
	Createat   int     `gorm:"column:CreateAt" json:"createAt"`
	IDproduct  int     `gorm:"column:IDProduct" json:"idProduct"`
	Product    Product `gorm:"joinForeignKey:IDProduct;foreignKey:ID;references:IDproduct" json:"productList"`
	IDcustomer int     `gorm:"column:IDCustomer" json:"idCustomer"`
	User       User    `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer" json:"userList"`
	Rate       int     `gorm:"column:Rate" json:"rate"`
}

// TableName get sql table name.获取数据库表名
func (m *Comment) TableName() string {
	return "Comment"
}

// CommentColumns get sql column name.获取数据库列名
var CommentColumns = struct {
	ID         string
	Content    string
	Createat   string
	IDproduct  string
	IDcustomer string
	Rate       string
}{
	ID:         "ID",
	Content:    "Content",
	Createat:   "CreateAt",
	IDproduct:  "IDProduct",
	IDcustomer: "IDCustomer",
	Rate:       "Rate",
}
