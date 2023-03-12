package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RankMgr struct {
	*_BaseMgr
}

// RankMgr open func
func RankMgr(db *gorm.DB) *_RankMgr {
	if db == nil {
		panic(fmt.Errorf("RankMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RankMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Rank"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RankMgr) GetTableName() string {
	return "Rank"
}

// Reset 重置gorm会话
func (obj *_RankMgr) Reset() *_RankMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RankMgr) Get() (result Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RankMgr) Gets() (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RankMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Rank{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_RankMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_RankMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithTotalspend TotalSpend获取
func (obj *_RankMgr) WithTotalspend(totalspend float32) Option {
	return optionFunc(func(o *options) { o.query["TotalSpend"] = totalspend })
}

// WithDiscount Discount获取
func (obj *_RankMgr) WithDiscount(discount float32) Option {
	return optionFunc(func(o *options) { o.query["Discount"] = discount })
}

// GetByOption 功能选项模式获取
func (obj *_RankMgr) GetByOption(opts ...Option) (result Rank, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RankMgr) GetByOptions(opts ...Option) (results []*Rank, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_RankMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Rank, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Rank{}).Where(options.query)
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
func (obj *_RankMgr) GetFromID(id int) (result Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`ID` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_RankMgr) GetBatchFromID(ids []int) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过Name获取内容
func (obj *_RankMgr) GetFromName(name string) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`Name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_RankMgr) GetBatchFromName(names []string) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`Name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTotalspend 通过TotalSpend获取内容
func (obj *_RankMgr) GetFromTotalspend(totalspend float32) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`TotalSpend` = ?", totalspend).Find(&results).Error

	return
}

// GetBatchFromTotalspend 批量查找
func (obj *_RankMgr) GetBatchFromTotalspend(totalspends []float32) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`TotalSpend` IN (?)", totalspends).Find(&results).Error

	return
}

// GetFromDiscount 通过Discount获取内容
func (obj *_RankMgr) GetFromDiscount(discount float32) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`Discount` = ?", discount).Find(&results).Error

	return
}

// GetBatchFromDiscount 批量查找
func (obj *_RankMgr) GetBatchFromDiscount(discounts []float32) (results []*Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`Discount` IN (?)", discounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RankMgr) FetchByPrimaryKey(id int) (result Rank, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rank{}).Where("`ID` = ?", id).First(&result).Error

	return
}
