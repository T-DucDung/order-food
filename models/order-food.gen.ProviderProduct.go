package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProviderproductMgr struct {
	*_BaseMgr
}

// ProviderproductMgr open func
func ProviderproductMgr(db *gorm.DB) *_ProviderproductMgr {
	if db == nil {
		panic(fmt.Errorf("ProviderproductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProviderproductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ProviderProduct"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProviderproductMgr) GetTableName() string {
	return "ProviderProduct"
}

// Reset 重置gorm会话
func (obj *_ProviderproductMgr) Reset() *_ProviderproductMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProviderproductMgr) Get() (result Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ProviderproductMgr) Gets() (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProviderproductMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIDproduct IDProduct获取
func (obj *_ProviderproductMgr) WithIDproduct(idproduct int) Option {
	return optionFunc(func(o *options) { o.query["IDProduct"] = idproduct })
}

// WithIDprovider IDProvider获取
func (obj *_ProviderproductMgr) WithIDprovider(idprovider int) Option {
	return optionFunc(func(o *options) { o.query["IDProvider"] = idprovider })
}

// WithPrice Price获取
func (obj *_ProviderproductMgr) WithPrice(price *float32) Option {
	return optionFunc(func(o *options) { o.query["Price"] = price })
}

// GetByOption 功能选项模式获取
func (obj *_ProviderproductMgr) GetByOption(opts ...Option) (result Providerproduct, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProviderproductMgr) GetByOptions(opts ...Option) (results []*Providerproduct, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_ProviderproductMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Providerproduct, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
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

// GetFromIDproduct 通过IDProduct获取内容
func (obj *_ProviderproductMgr) GetFromIDproduct(idproduct int) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDproduct 批量查找
func (obj *_ProviderproductMgr) GetBatchFromIDproduct(idproducts []int) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`IDProduct` IN (?)", idproducts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDprovider 通过IDProvider获取内容
func (obj *_ProviderproductMgr) GetFromIDprovider(idprovider int) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`IDProvider` = ?", idprovider).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDprovider 批量查找
func (obj *_ProviderproductMgr) GetBatchFromIDprovider(idproviders []int) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`IDProvider` IN (?)", idproviders).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPrice 通过Price获取内容
func (obj *_ProviderproductMgr) GetFromPrice(price *float32) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`Price` = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPrice 批量查找
func (obj *_ProviderproductMgr) GetBatchFromPrice(prices []*float32) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`Price` IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
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
func (obj *_ProviderproductMgr) FetchByPrimaryKey(idproduct int, idprovider int) (result Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`IDProduct` = ? AND `IDProvider` = ?", idproduct, idprovider).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByProviderproductFk1  获取多个内容
func (obj *_ProviderproductMgr) FetchIndexByProviderproductFk1(idprovider int) (results []*Providerproduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Providerproduct{}).Where("`IDProvider` = ?", idprovider).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
