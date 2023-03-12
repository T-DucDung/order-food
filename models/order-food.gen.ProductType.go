package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProducttypeMgr struct {
	*_BaseMgr
}

// ProducttypeMgr open func
func ProducttypeMgr(db *gorm.DB) *_ProducttypeMgr {
	if db == nil {
		panic(fmt.Errorf("ProducttypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProducttypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ProductType"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProducttypeMgr) GetTableName() string {
	return "ProductType"
}

// Reset 重置gorm会话
func (obj *_ProducttypeMgr) Reset() *_ProducttypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProducttypeMgr) Get() (result Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProducttypeMgr) Gets() (results []*Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProducttypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Producttype{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_ProducttypeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_ProducttypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_ProducttypeMgr) GetByOption(opts ...Option) (result Producttype, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProducttypeMgr) GetByOptions(opts ...Option) (results []*Producttype, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ProducttypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Producttype, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where(options.query)
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
func (obj *_ProducttypeMgr) GetFromID(id int) (result Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where("`ID` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ProducttypeMgr) GetBatchFromID(ids []int) (results []*Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过Name获取内容
func (obj *_ProducttypeMgr) GetFromName(name string) (results []*Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where("`Name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_ProducttypeMgr) GetBatchFromName(names []string) (results []*Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where("`Name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProducttypeMgr) FetchByPrimaryKey(id int) (result Producttype, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Producttype{}).Where("`ID` = ?", id).First(&result).Error

	return
}
