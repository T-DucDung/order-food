package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProviderMgr struct {
	*_BaseMgr
}

// ProviderMgr open func
func ProviderMgr(db *gorm.DB) *_ProviderMgr {
	if db == nil {
		panic(fmt.Errorf("ProviderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProviderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Provider"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProviderMgr) GetTableName() string {
	return "Provider"
}

// Reset 重置gorm会话
func (obj *_ProviderMgr) Reset() *_ProviderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProviderMgr) Get() (result Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProviderMgr) Gets() (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProviderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Provider{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_ProviderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_ProviderMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithPhone Phone获取
func (obj *_ProviderMgr) WithPhone(phone *int) Option {
	return optionFunc(func(o *options) { o.query["Phone"] = phone })
}

// WithAddress Address获取
func (obj *_ProviderMgr) WithAddress(address *string) Option {
	return optionFunc(func(o *options) { o.query["Address"] = address })
}

// WithEmail Email获取
func (obj *_ProviderMgr) WithEmail(email *string) Option {
	return optionFunc(func(o *options) { o.query["Email"] = email })
}

// GetByOption 功能选项模式获取
func (obj *_ProviderMgr) GetByOption(opts ...Option) (result Provider, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProviderMgr) GetByOptions(opts ...Option) (results []*Provider, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ProviderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Provider, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Provider{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过ID获取内容
func (obj *_ProviderMgr) GetFromID(id int) (result Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`ID` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ProviderMgr) GetBatchFromID(ids []int) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过Name获取内容
func (obj *_ProviderMgr) GetFromName(name string) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_ProviderMgr) GetBatchFromName(names []string) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过Phone获取内容
func (obj *_ProviderMgr) GetFromPhone(phone *int) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_ProviderMgr) GetBatchFromPhone(phones []*int) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddress 通过Address获取内容
func (obj *_ProviderMgr) GetFromAddress(address *string) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找
func (obj *_ProviderMgr) GetBatchFromAddress(addresss []*string) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromEmail 通过Email获取内容
func (obj *_ProviderMgr) GetFromEmail(email *string) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_ProviderMgr) GetBatchFromEmail(emails []*string) (results []*Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`Email` IN (?)", emails).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProviderMgr) FetchByPrimaryKey(id int) (result Provider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Provider{}).Where("`ID` = ?", id).First(&result).Error

	return
}
