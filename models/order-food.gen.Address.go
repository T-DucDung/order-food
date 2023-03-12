package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AddressMgr struct {
	*_BaseMgr
}

// AddressMgr open func
func AddressMgr(db *gorm.DB) *_AddressMgr {
	if db == nil {
		panic(fmt.Errorf("AddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Address"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AddressMgr) GetTableName() string {
	return "Address"
}

// Reset 重置gorm会话
func (obj *_AddressMgr) Reset() *_AddressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AddressMgr) Get() (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).First(&result).Error
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
func (obj *_AddressMgr) Gets() (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Find(&results).Error
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
func (obj *_AddressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Address{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_AddressMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_AddressMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithPhone Phone获取
func (obj *_AddressMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["Phone"] = phone })
}

// WithAddress Address获取
func (obj *_AddressMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["Address"] = address })
}

// WithIDcustomer IDCustomer获取
func (obj *_AddressMgr) WithIDcustomer(idcustomer int) Option {
	return optionFunc(func(o *options) { o.query["IDCustomer"] = idcustomer })
}

// GetByOption 功能选项模式获取
func (obj *_AddressMgr) GetByOption(opts ...Option) (result Address, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where(options.query).First(&result).Error
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
func (obj *_AddressMgr) GetByOptions(opts ...Option) (results []*Address, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where(options.query).Find(&results).Error
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
func (obj *_AddressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Address, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Address{}).Where(options.query)
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
func (obj *_AddressMgr) GetFromID(id int) (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`ID` = ?", id).First(&result).Error
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
func (obj *_AddressMgr) GetBatchFromID(ids []int) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`ID` IN (?)", ids).Find(&results).Error
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

// GetFromName 通过Name获取内容
func (obj *_AddressMgr) GetFromName(name string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`Name` = ?", name).Find(&results).Error
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

// GetBatchFromName 批量查找
func (obj *_AddressMgr) GetBatchFromName(names []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`Name` IN (?)", names).Find(&results).Error
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

// GetFromPhone 通过Phone获取内容
func (obj *_AddressMgr) GetFromPhone(phone string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`Phone` = ?", phone).Find(&results).Error
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

// GetBatchFromPhone 批量查找
func (obj *_AddressMgr) GetBatchFromPhone(phones []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`Phone` IN (?)", phones).Find(&results).Error
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

// GetFromAddress 通过Address获取内容
func (obj *_AddressMgr) GetFromAddress(address string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`Address` = ?", address).Find(&results).Error
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

// GetBatchFromAddress 批量查找
func (obj *_AddressMgr) GetBatchFromAddress(addresss []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`Address` IN (?)", addresss).Find(&results).Error
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
func (obj *_AddressMgr) GetFromIDcustomer(idcustomer int) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
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
func (obj *_AddressMgr) GetBatchFromIDcustomer(idcustomers []int) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`IDCustomer` IN (?)", idcustomers).Find(&results).Error
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
func (obj *_AddressMgr) FetchByPrimaryKey(id int) (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByAddressFk  获取多个内容
func (obj *_AddressMgr) FetchIndexByAddressFk(idcustomer int) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Address{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
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
