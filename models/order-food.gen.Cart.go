package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CartMgr struct {
	*_BaseMgr
}

// CartMgr open func
func CartMgr(db *gorm.DB) *_CartMgr {
	if db == nil {
		panic(fmt.Errorf("CartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Cart"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartMgr) GetTableName() string {
	return "Cart"
}

// Reset 重置gorm会话
func (obj *_CartMgr) Reset() *_CartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CartMgr) Get() (result Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_CartMgr) Gets() (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Cart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_CartMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithIDcustomer IDCustomer获取
func (obj *_CartMgr) WithIDcustomer(idcustomer int) Option {
	return optionFunc(func(o *options) { o.query["IDCustomer"] = idcustomer })
}

// WithTotal Total获取
func (obj *_CartMgr) WithTotal(total float32) Option {
	return optionFunc(func(o *options) { o.query["Total"] = total })
}

// GetByOption 功能选项模式获取
func (obj *_CartMgr) GetByOption(opts ...Option) (result Cart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CartMgr) GetByOptions(opts ...Option) (results []*Cart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_CartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Cart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Cart{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
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

// GetFromID 通过ID获取内容
func (obj *_CartMgr) GetFromID(id int) (result Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_CartMgr) GetBatchFromID(ids []int) (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDcustomer 通过IDCustomer获取内容
func (obj *_CartMgr) GetFromIDcustomer(idcustomer int) (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDcustomer 批量查找
func (obj *_CartMgr) GetBatchFromIDcustomer(idcustomers []int) (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`IDCustomer` IN (?)", idcustomers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTotal 通过Total获取内容
func (obj *_CartMgr) GetFromTotal(total float32) (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`Total` = ?", total).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTotal 批量查找
func (obj *_CartMgr) GetBatchFromTotal(totals []float32) (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`Total` IN (?)", totals).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
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
func (obj *_CartMgr) FetchByPrimaryKey(id int) (result Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByCartFk  获取多个内容
func (obj *_CartMgr) FetchIndexByCartFk(idcustomer int) (results []*Cart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Cart{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
