package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _InventoryreceivingvoucherMgr struct {
	*_BaseMgr
}

// InventoryreceivingvoucherMgr open func
func InventoryreceivingvoucherMgr(db *gorm.DB) *_InventoryreceivingvoucherMgr {
	if db == nil {
		panic(fmt.Errorf("InventoryreceivingvoucherMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InventoryreceivingvoucherMgr{_BaseMgr: &_BaseMgr{DB: db.Table("InventoryReceivingVoucher"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InventoryreceivingvoucherMgr) GetTableName() string {
	return "InventoryReceivingVoucher"
}

// Reset 重置gorm会话
func (obj *_InventoryreceivingvoucherMgr) Reset() *_InventoryreceivingvoucherMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_InventoryreceivingvoucherMgr) Get() (result Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_InventoryreceivingvoucherMgr) Gets() (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_InventoryreceivingvoucherMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_InventoryreceivingvoucherMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithCreateat CreateAt获取
func (obj *_InventoryreceivingvoucherMgr) WithCreateat(createat int) Option {
	return optionFunc(func(o *options) { o.query["CreateAt"] = createat })
}

// WithStatus Status获取
func (obj *_InventoryreceivingvoucherMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["Status"] = status })
}

// WithTotalprice TotalPrice获取
func (obj *_InventoryreceivingvoucherMgr) WithTotalprice(totalprice float32) Option {
	return optionFunc(func(o *options) { o.query["TotalPrice"] = totalprice })
}

// WithIDprovider IDProvider获取
func (obj *_InventoryreceivingvoucherMgr) WithIDprovider(idprovider int) Option {
	return optionFunc(func(o *options) { o.query["IDProvider"] = idprovider })
}

// WithIDcreater IDCreater获取
func (obj *_InventoryreceivingvoucherMgr) WithIDcreater(idcreater int) Option {
	return optionFunc(func(o *options) { o.query["IDCreater"] = idcreater })
}

// WithIDimporter IDImporter获取
func (obj *_InventoryreceivingvoucherMgr) WithIDimporter(idimporter int) Option {
	return optionFunc(func(o *options) { o.query["IDImporter"] = idimporter })
}

// GetByOption 功能选项模式获取
func (obj *_InventoryreceivingvoucherMgr) GetByOption(opts ...Option) (result Inventoryreceivingvoucher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_InventoryreceivingvoucherMgr) GetByOptions(opts ...Option) (results []*Inventoryreceivingvoucher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_InventoryreceivingvoucherMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Inventoryreceivingvoucher, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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

// GetFromID 通过ID获取内容
func (obj *_InventoryreceivingvoucherMgr) GetFromID(id int) (result Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromID(ids []int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateat 通过CreateAt获取内容
func (obj *_InventoryreceivingvoucherMgr) GetFromCreateat(createat int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`CreateAt` = ?", createat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateat 批量查找
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromCreateat(createats []int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`CreateAt` IN (?)", createats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过Status获取内容
func (obj *_InventoryreceivingvoucherMgr) GetFromStatus(status int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`Status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromStatus(statuss []int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`Status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTotalprice 通过TotalPrice获取内容
func (obj *_InventoryreceivingvoucherMgr) GetFromTotalprice(totalprice float32) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`TotalPrice` = ?", totalprice).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTotalprice 批量查找
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromTotalprice(totalprices []float32) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`TotalPrice` IN (?)", totalprices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_InventoryreceivingvoucherMgr) GetFromIDprovider(idprovider int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDProvider` = ?", idprovider).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromIDprovider(idproviders []int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDProvider` IN (?)", idproviders).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDcreater 通过IDCreater获取内容
func (obj *_InventoryreceivingvoucherMgr) GetFromIDcreater(idcreater int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDCreater` = ?", idcreater).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDcreater 批量查找
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromIDcreater(idcreaters []int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDCreater` IN (?)", idcreaters).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDimporter 通过IDImporter获取内容
func (obj *_InventoryreceivingvoucherMgr) GetFromIDimporter(idimporter int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDImporter` = ?", idimporter).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDimporter 批量查找
func (obj *_InventoryreceivingvoucherMgr) GetBatchFromIDimporter(idimporters []int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDImporter` IN (?)", idimporters).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_InventoryreceivingvoucherMgr) FetchByPrimaryKey(id int) (result Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Provider").Where("ID = ?", result.IDprovider).Find(&result.Provider).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByInventoryreceivingvoucherFk  获取多个内容
func (obj *_InventoryreceivingvoucherMgr) FetchIndexByInventoryreceivingvoucherFk(idprovider int) (results []*Inventoryreceivingvoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Inventoryreceivingvoucher{}).Where("`IDProvider` = ?", idprovider).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Provider").Where("ID = ?", results[i].IDprovider).Find(&results[i].Provider).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
