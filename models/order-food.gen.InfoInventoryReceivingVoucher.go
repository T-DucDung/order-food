package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _InfoinventoryreceivingvoucherMgr struct {
	*_BaseMgr
}

// InfoinventoryreceivingvoucherMgr open func
func InfoinventoryreceivingvoucherMgr(db *gorm.DB) *_InfoinventoryreceivingvoucherMgr {
	if db == nil {
		panic(fmt.Errorf("InfoinventoryreceivingvoucherMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoinventoryreceivingvoucherMgr{_BaseMgr: &_BaseMgr{DB: db.Table("InfoInventoryReceivingVoucher"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InfoinventoryreceivingvoucherMgr) GetTableName() string {
	return "InfoInventoryReceivingVoucher"
}

// Reset 重置gorm会话
func (obj *_InfoinventoryreceivingvoucherMgr) Reset() *_InfoinventoryreceivingvoucherMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_InfoinventoryreceivingvoucherMgr) Get() (result Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", result.IDvoucher).Find(&result.Inventoryreceivingvoucher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_InfoinventoryreceivingvoucherMgr) Gets() (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_InfoinventoryreceivingvoucherMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIDproduct IDProduct获取
func (obj *_InfoinventoryreceivingvoucherMgr) WithIDproduct(idproduct int) Option {
	return optionFunc(func(o *options) { o.query["IDProduct"] = idproduct })
}

// WithIDvoucher IDVoucher获取
func (obj *_InfoinventoryreceivingvoucherMgr) WithIDvoucher(idvoucher int) Option {
	return optionFunc(func(o *options) { o.query["IDVoucher"] = idvoucher })
}

// WithTotal Total获取
func (obj *_InfoinventoryreceivingvoucherMgr) WithTotal(total int) Option {
	return optionFunc(func(o *options) { o.query["Total"] = total })
}

// WithPrice Price获取
func (obj *_InfoinventoryreceivingvoucherMgr) WithPrice(price float32) Option {
	return optionFunc(func(o *options) { o.query["Price"] = price })
}

// GetByOption 功能选项模式获取
func (obj *_InfoinventoryreceivingvoucherMgr) GetByOption(opts ...Option) (result Infoinventoryreceivingvoucher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", result.IDvoucher).Find(&result.Inventoryreceivingvoucher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_InfoinventoryreceivingvoucherMgr) GetByOptions(opts ...Option) (results []*Infoinventoryreceivingvoucher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_InfoinventoryreceivingvoucherMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Infoinventoryreceivingvoucher, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where(options.query)
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
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
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
func (obj *_InfoinventoryreceivingvoucherMgr) GetFromIDproduct(idproduct int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDproduct 批量查找
func (obj *_InfoinventoryreceivingvoucherMgr) GetBatchFromIDproduct(idproducts []int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`IDProduct` IN (?)", idproducts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDvoucher 通过IDVoucher获取内容
func (obj *_InfoinventoryreceivingvoucherMgr) GetFromIDvoucher(idvoucher int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`IDVoucher` = ?", idvoucher).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDvoucher 批量查找
func (obj *_InfoinventoryreceivingvoucherMgr) GetBatchFromIDvoucher(idvouchers []int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`IDVoucher` IN (?)", idvouchers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTotal 通过Total获取内容
func (obj *_InfoinventoryreceivingvoucherMgr) GetFromTotal(total int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`Total` = ?", total).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTotal 批量查找
func (obj *_InfoinventoryreceivingvoucherMgr) GetBatchFromTotal(totals []int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`Total` IN (?)", totals).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPrice 通过Price获取内容
func (obj *_InfoinventoryreceivingvoucherMgr) GetFromPrice(price float32) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`Price` = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPrice 批量查找
func (obj *_InfoinventoryreceivingvoucherMgr) GetBatchFromPrice(prices []float32) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`Price` IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
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
func (obj *_InfoinventoryreceivingvoucherMgr) FetchByPrimaryKey(idproduct int, idvoucher int) (result Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`IDProduct` = ? AND `IDVoucher` = ?", idproduct, idvoucher).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", result.IDvoucher).Find(&result.Inventoryreceivingvoucher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByInfoinventoryreceivingvoucherFk  获取多个内容
func (obj *_InfoinventoryreceivingvoucherMgr) FetchIndexByInfoinventoryreceivingvoucherFk(idvoucher int) (results []*Infoinventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Infoinventoryreceivingvoucher{}).Where("`IDVoucher` = ?", idvoucher).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("InventoryReceivingVoucher").Where("ID = ?", results[i].IDvoucher).Find(&results[i].Inventoryreceivingvoucher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
