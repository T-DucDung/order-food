package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _InfocartMgr struct {
	*_BaseMgr
}

// InfocartMgr open func
func InfocartMgr(db *gorm.DB) *_InfocartMgr {
	if db == nil {
		panic(fmt.Errorf("InfocartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfocartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("InfoCart"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InfocartMgr) GetTableName() string {
	return "InfoCart"
}

// Reset 重置gorm会话
func (obj *_InfocartMgr) Reset() *_InfocartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_InfocartMgr) Get() (result Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Cart").Where("ID = ?", result.IDcart).Find(&result.Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) Gets() (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Infocart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIDcart IDCart获取
func (obj *_InfocartMgr) WithIDcart(idcart int) Option {
	return optionFunc(func(o *options) { o.query["IDCart"] = idcart })
}

// WithIDproduct IDProduct获取
func (obj *_InfocartMgr) WithIDproduct(idproduct int) Option {
	return optionFunc(func(o *options) { o.query["IDProduct"] = idproduct })
}

// WithQuantity Quantity获取
func (obj *_InfocartMgr) WithQuantity(quantity float32) Option {
	return optionFunc(func(o *options) { o.query["Quantity"] = quantity })
}

// WithPrice Price获取
func (obj *_InfocartMgr) WithPrice(price float32) Option {
	return optionFunc(func(o *options) { o.query["Price"] = price })
}

// WithPricesale PriceSale获取
func (obj *_InfocartMgr) WithPricesale(pricesale *float32) Option {
	return optionFunc(func(o *options) { o.query["PriceSale"] = pricesale })
}

// GetByOption 功能选项模式获取
func (obj *_InfocartMgr) GetByOption(opts ...Option) (result Infocart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Cart").Where("ID = ?", result.IDcart).Find(&result.Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetByOptions(opts ...Option) (results []*Infocart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Infocart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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

// GetFromIDcart 通过IDCart获取内容
func (obj *_InfocartMgr) GetFromIDcart(idcart int) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`IDCart` = ?", idcart).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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

// GetBatchFromIDcart 批量查找
func (obj *_InfocartMgr) GetBatchFromIDcart(idcarts []int) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`IDCart` IN (?)", idcarts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetFromIDproduct(idproduct int) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetBatchFromIDproduct(idproducts []int) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`IDProduct` IN (?)", idproducts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetFromQuantity(quantity float32) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`Quantity` = ?", quantity).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetBatchFromQuantity(quantitys []float32) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`Quantity` IN (?)", quantitys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetFromPrice(price float32) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`Price` = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetBatchFromPrice(prices []float32) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`Price` IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetFromPricesale(pricesale *float32) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`PriceSale` = ?", pricesale).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) GetBatchFromPricesale(pricesales []*float32) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`PriceSale` IN (?)", pricesales).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) FetchByPrimaryKey(idcart int, idproduct int) (result Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`IDCart` = ? AND `IDProduct` = ?", idcart, idproduct).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Cart").Where("ID = ?", result.IDcart).Find(&result.Cart).Error; err != nil { //
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
func (obj *_InfocartMgr) FetchIndexByInfocartFk1(idproduct int) (results []*Infocart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infocart{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Cart").Where("ID = ?", results[i].IDcart).Find(&results[i].Cart).Error; err != nil { //
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
