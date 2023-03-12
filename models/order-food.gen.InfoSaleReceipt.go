package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _InfosalereceiptMgr struct {
	*_BaseMgr
}

// InfosalereceiptMgr open func
func InfosalereceiptMgr(db *gorm.DB) *_InfosalereceiptMgr {
	if db == nil {
		panic(fmt.Errorf("InfosalereceiptMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfosalereceiptMgr{_BaseMgr: &_BaseMgr{DB: db.Table("InfoSaleReceipt"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InfosalereceiptMgr) GetTableName() string {
	return "InfoSaleReceipt"
}

// Reset 重置gorm会话
func (obj *_InfosalereceiptMgr) Reset() *_InfosalereceiptMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_InfosalereceiptMgr) Get() (result Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", result.IDsalereceipt).Find(&result.Salereceipt).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_InfosalereceiptMgr) Gets() (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_InfosalereceiptMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIDsalereceipt IDSaleReceipt获取
func (obj *_InfosalereceiptMgr) WithIDsalereceipt(idsalereceipt int) Option {
	return optionFunc(func(o *options) { o.query["IDSaleReceipt"] = idsalereceipt })
}

// WithIDproduct IDProduct获取
func (obj *_InfosalereceiptMgr) WithIDproduct(idproduct int) Option {
	return optionFunc(func(o *options) { o.query["IDProduct"] = idproduct })
}

// WithQuantity Quantity获取
func (obj *_InfosalereceiptMgr) WithQuantity(quantity float32) Option {
	return optionFunc(func(o *options) { o.query["Quantity"] = quantity })
}

// WithPrice Price获取
func (obj *_InfosalereceiptMgr) WithPrice(price float32) Option {
	return optionFunc(func(o *options) { o.query["Price"] = price })
}

// WithPricesale PriceSale获取
func (obj *_InfosalereceiptMgr) WithPricesale(pricesale *float32) Option {
	return optionFunc(func(o *options) { o.query["PriceSale"] = pricesale })
}

// GetByOption 功能选项模式获取
func (obj *_InfosalereceiptMgr) GetByOption(opts ...Option) (result Infosalereceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", result.IDsalereceipt).Find(&result.Salereceipt).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_InfosalereceiptMgr) GetByOptions(opts ...Option) (results []*Infosalereceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_InfosalereceiptMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Infosalereceipt, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDsalereceipt 通过IDSaleReceipt获取内容
func (obj *_InfosalereceiptMgr) GetFromIDsalereceipt(idsalereceipt int) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`IDSaleReceipt` = ?", idsalereceipt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDsalereceipt 批量查找
func (obj *_InfosalereceiptMgr) GetBatchFromIDsalereceipt(idsalereceipts []int) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`IDSaleReceipt` IN (?)", idsalereceipts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDproduct 通过IDProduct获取内容
func (obj *_InfosalereceiptMgr) GetFromIDproduct(idproduct int) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDproduct 批量查找
func (obj *_InfosalereceiptMgr) GetBatchFromIDproduct(idproducts []int) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`IDProduct` IN (?)", idproducts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromQuantity 通过Quantity获取内容
func (obj *_InfosalereceiptMgr) GetFromQuantity(quantity float32) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`Quantity` = ?", quantity).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromQuantity 批量查找
func (obj *_InfosalereceiptMgr) GetBatchFromQuantity(quantitys []float32) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`Quantity` IN (?)", quantitys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPrice 通过Price获取内容
func (obj *_InfosalereceiptMgr) GetFromPrice(price float32) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`Price` = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPrice 批量查找
func (obj *_InfosalereceiptMgr) GetBatchFromPrice(prices []float32) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`Price` IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPricesale 通过PriceSale获取内容
func (obj *_InfosalereceiptMgr) GetFromPricesale(pricesale *float32) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`PriceSale` = ?", pricesale).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPricesale 批量查找
func (obj *_InfosalereceiptMgr) GetBatchFromPricesale(pricesales []*float32) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`PriceSale` IN (?)", pricesales).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_InfosalereceiptMgr) FetchByPrimaryKey(idsalereceipt int, idproduct int) (result Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`IDSaleReceipt` = ? AND `IDProduct` = ?", idsalereceipt, idproduct).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", result.IDsalereceipt).Find(&result.Salereceipt).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByInfocartFk1  获取多个内容
func (obj *_InfosalereceiptMgr) FetchIndexByInfocartFk1(idproduct int) (results []*Infosalereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infosalereceipt{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("SaleReceipt").Where("ID = ?", results[i].IDsalereceipt).Find(&results[i].Salereceipt).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
