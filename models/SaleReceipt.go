package models

// Salereceipt [...]
type Salereceipt struct {
	ID           int     `gorm:"primaryKey;column:ID"`
	Createat     int     `gorm:"column:CreateAt"`
	Deliverydate int     `gorm:"column:DeliveryDate"`
	Ispayment    bool    `gorm:"column:IsPayment"`
	Note         string  `gorm:"column:Note"`
	Typepayment  string  `gorm:"column:TypePayment"`
	Status       int     `gorm:"column:Status"`
	Total        float32 `gorm:"column:Total"`
	IDcustomer   int     `gorm:"column:IDCustomer"`
	User         User    `gorm:"joinForeignKey:IDCustomer;foreignKey:ID;references:IDcustomer"`
	IDdeliver    int     `gorm:"column:IDDeliver"`
	Rankdiscount float32 `gorm:"column:RankDiscount"`
	IDreceiver   int     `gorm:"column:IDReceiver"`
	Address      Address `gorm:"joinForeignKey:IDReceiver;foreignKey:ID;references:IDreceiver"`
}

// TableName get sql table name.获取数据库表名
func (m *Salereceipt) TableName() string {
	return "SaleReceipt"
}
