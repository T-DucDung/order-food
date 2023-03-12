package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProductMgr struct {
	*_BaseMgr
}

// ProductMgr open func
func ProductMgr(db *gorm.DB) *_ProductMgr {
	if db == nil {
		panic(fmt.Errorf("ProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductMgr) GetTableName() string {
	return "Product"
}

// Reset 重置gorm会话
func (obj *_ProductMgr) Reset() *_ProductMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductMgr) Get() (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ProductType").Where("ID = ?", result.IDtype).Find(&result.Producttype).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ProductMgr) Gets() (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Product{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_ProductMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_ProductMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithImg Img获取
func (obj *_ProductMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["Img"] = img })
}

// WithPrice Price获取
func (obj *_ProductMgr) WithPrice(price float32) Option {
	return optionFunc(func(o *options) { o.query["Price"] = price })
}

// WithUnit unit获取
func (obj *_ProductMgr) WithUnit(unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = unit })
}

// WithRemain Remain获取
func (obj *_ProductMgr) WithRemain(remain float32) Option {
	return optionFunc(func(o *options) { o.query["Remain"] = remain })
}

// WithSold Sold获取
func (obj *_ProductMgr) WithSold(sold float32) Option {
	return optionFunc(func(o *options) { o.query["Sold"] = sold })
}

// WithDescription Description获取
func (obj *_ProductMgr) WithDescription(description *string) Option {
	return optionFunc(func(o *options) { o.query["Description"] = description })
}

// WithRateavg RateAVG获取
func (obj *_ProductMgr) WithRateavg(rateavg int) Option {
	return optionFunc(func(o *options) { o.query["RateAVG"] = rateavg })
}

// WithIDtype IDType获取
func (obj *_ProductMgr) WithIDtype(idtype int) Option {
	return optionFunc(func(o *options) { o.query["IDType"] = idtype })
}

// GetByOption 功能选项模式获取
func (obj *_ProductMgr) GetByOption(opts ...Option) (result Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ProductType").Where("ID = ?", result.IDtype).Find(&result.Producttype).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductMgr) GetByOptions(opts ...Option) (results []*Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_ProductMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Product, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
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
func (obj *_ProductMgr) GetFromID(id int) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ProductType").Where("ID = ?", result.IDtype).Find(&result.Producttype).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_ProductMgr) GetBatchFromID(ids []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过Name获取内容
func (obj *_ProductMgr) GetFromName(name string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找
func (obj *_ProductMgr) GetBatchFromName(names []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromImg 通过Img获取内容
func (obj *_ProductMgr) GetFromImg(img string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Img` = ?", img).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromImg 批量查找
func (obj *_ProductMgr) GetBatchFromImg(imgs []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Img` IN (?)", imgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPrice 通过Price获取内容
func (obj *_ProductMgr) GetFromPrice(price float32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Price` = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPrice 批量查找
func (obj *_ProductMgr) GetBatchFromPrice(prices []float32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Price` IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUnit 通过unit获取内容
func (obj *_ProductMgr) GetFromUnit(unit string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`unit` = ?", unit).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUnit 批量查找
func (obj *_ProductMgr) GetBatchFromUnit(units []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`unit` IN (?)", units).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRemain 通过Remain获取内容
func (obj *_ProductMgr) GetFromRemain(remain float32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Remain` = ?", remain).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRemain 批量查找
func (obj *_ProductMgr) GetBatchFromRemain(remains []float32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Remain` IN (?)", remains).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSold 通过Sold获取内容
func (obj *_ProductMgr) GetFromSold(sold float32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Sold` = ?", sold).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSold 批量查找
func (obj *_ProductMgr) GetBatchFromSold(solds []float32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Sold` IN (?)", solds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDescription 通过Description获取内容
func (obj *_ProductMgr) GetFromDescription(description *string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Description` = ?", description).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDescription 批量查找
func (obj *_ProductMgr) GetBatchFromDescription(descriptions []*string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`Description` IN (?)", descriptions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRateavg 通过RateAVG获取内容
func (obj *_ProductMgr) GetFromRateavg(rateavg int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`RateAVG` = ?", rateavg).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRateavg 批量查找
func (obj *_ProductMgr) GetBatchFromRateavg(rateavgs []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`RateAVG` IN (?)", rateavgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDtype 通过IDType获取内容
func (obj *_ProductMgr) GetFromIDtype(idtype int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`IDType` = ?", idtype).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDtype 批量查找
func (obj *_ProductMgr) GetBatchFromIDtype(idtypes []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`IDType` IN (?)", idtypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
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
func (obj *_ProductMgr) FetchByPrimaryKey(id int) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ProductType").Where("ID = ?", result.IDtype).Find(&result.Producttype).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByProductFk  获取多个内容
func (obj *_ProductMgr) FetchIndexByProductFk(idtype int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`IDType` = ?", idtype).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ProductType").Where("ID = ?", results[i].IDtype).Find(&results[i].Producttype).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
