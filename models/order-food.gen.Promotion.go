package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PromotionMgr struct {
	*_BaseMgr
}

// PromotionMgr open func
func PromotionMgr(db *gorm.DB) *_PromotionMgr {
	if db == nil {
		panic(fmt.Errorf("PromotionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PromotionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Promotion"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PromotionMgr) GetTableName() string {
	return "Promotion"
}

// Reset 重置gorm会话
func (obj *_PromotionMgr) Reset() *_PromotionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PromotionMgr) Get() (result Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.Productid).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_PromotionMgr) Gets() (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PromotionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Promotion{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_PromotionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithCreateat CreateAt获取
func (obj *_PromotionMgr) WithCreateat(createat int) Option {
	return optionFunc(func(o *options) { o.query["CreateAt"] = createat })
}

// WithStartat StartAt获取
func (obj *_PromotionMgr) WithStartat(startat int) Option {
	return optionFunc(func(o *options) { o.query["StartAt"] = startat })
}

// WithEndat EndAt获取
func (obj *_PromotionMgr) WithEndat(endat int) Option {
	return optionFunc(func(o *options) { o.query["EndAt"] = endat })
}

// WithPrice Price获取
func (obj *_PromotionMgr) WithPrice(price float32) Option {
	return optionFunc(func(o *options) { o.query["Price"] = price })
}

// WithProductid ProductID获取
func (obj *_PromotionMgr) WithProductid(productid int) Option {
	return optionFunc(func(o *options) { o.query["ProductID"] = productid })
}

// GetByOption 功能选项模式获取
func (obj *_PromotionMgr) GetByOption(opts ...Option) (result Promotion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.Productid).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PromotionMgr) GetByOptions(opts ...Option) (results []*Promotion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_PromotionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Promotion, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
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
func (obj *_PromotionMgr) GetFromID(id int) (result Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.Productid).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_PromotionMgr) GetBatchFromID(ids []int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateat 通过CreateAt获取内容
func (obj *_PromotionMgr) GetFromCreateat(createat int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`CreateAt` = ?", createat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateat 批量查找
func (obj *_PromotionMgr) GetBatchFromCreateat(createats []int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`CreateAt` IN (?)", createats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStartat 通过StartAt获取内容
func (obj *_PromotionMgr) GetFromStartat(startat int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`StartAt` = ?", startat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStartat 批量查找
func (obj *_PromotionMgr) GetBatchFromStartat(startats []int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`StartAt` IN (?)", startats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEndat 通过EndAt获取内容
func (obj *_PromotionMgr) GetFromEndat(endat int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`EndAt` = ?", endat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEndat 批量查找
func (obj *_PromotionMgr) GetBatchFromEndat(endats []int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`EndAt` IN (?)", endats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPrice 通过Price获取内容
func (obj *_PromotionMgr) GetFromPrice(price float32) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`Price` = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPrice 批量查找
func (obj *_PromotionMgr) GetBatchFromPrice(prices []float32) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`Price` IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProductid 通过ProductID获取内容
func (obj *_PromotionMgr) GetFromProductid(productid int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`ProductID` = ?", productid).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProductid 批量查找
func (obj *_PromotionMgr) GetBatchFromProductid(productids []int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`ProductID` IN (?)", productids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
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
func (obj *_PromotionMgr) FetchByPrimaryKey(id int) (result Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.Productid).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexBySaleinfoFk  获取多个内容
func (obj *_PromotionMgr) FetchIndexBySaleinfoFk(productid int) (results []*Promotion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Promotion{}).Where("`ProductID` = ?", productid).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].Productid).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
