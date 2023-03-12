package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SalereceiptMgr struct {
	*_BaseMgr
}

// SalereceiptMgr open func
func SalereceiptMgr(db *gorm.DB) *_SalereceiptMgr {
	if db == nil {
		panic(fmt.Errorf("SalereceiptMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SalereceiptMgr{_BaseMgr: &_BaseMgr{DB: db.Table("SaleReceipt"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SalereceiptMgr) GetTableName() string {
	return "SaleReceipt"
}

// Reset 重置gorm会话
func (obj *_SalereceiptMgr) Reset() *_SalereceiptMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SalereceiptMgr) Get() (result Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Address").Where("ID = ?", result.IDreceiver).Find(&result.Address).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_SalereceiptMgr) Gets() (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SalereceiptMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_SalereceiptMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithCreateat CreateAt获取
func (obj *_SalereceiptMgr) WithCreateat(createat int) Option {
	return optionFunc(func(o *options) { o.query["CreateAt"] = createat })
}

// WithDeliverydate DeliveryDate获取
func (obj *_SalereceiptMgr) WithDeliverydate(deliverydate *int) Option {
	return optionFunc(func(o *options) { o.query["DeliveryDate"] = deliverydate })
}

// WithIspayment IsPayment获取
func (obj *_SalereceiptMgr) WithIspayment(ispayment bool) Option {
	return optionFunc(func(o *options) { o.query["IsPayment"] = ispayment })
}

// WithNote Note获取
func (obj *_SalereceiptMgr) WithNote(note *string) Option {
	return optionFunc(func(o *options) { o.query["Note"] = note })
}

// WithTypepayment TypePayment获取
func (obj *_SalereceiptMgr) WithTypepayment(typepayment *string) Option {
	return optionFunc(func(o *options) { o.query["TypePayment"] = typepayment })
}

// WithStatus Status获取
func (obj *_SalereceiptMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["Status"] = status })
}

// WithTotal Total获取
func (obj *_SalereceiptMgr) WithTotal(total float32) Option {
	return optionFunc(func(o *options) { o.query["Total"] = total })
}

// WithIDcustomer IDCustomer获取
func (obj *_SalereceiptMgr) WithIDcustomer(idcustomer int) Option {
	return optionFunc(func(o *options) { o.query["IDCustomer"] = idcustomer })
}

// WithIDdeliver IDDeliver获取
func (obj *_SalereceiptMgr) WithIDdeliver(iddeliver int) Option {
	return optionFunc(func(o *options) { o.query["IDDeliver"] = iddeliver })
}

// WithRankdiscount RankDiscount获取
func (obj *_SalereceiptMgr) WithRankdiscount(rankdiscount float32) Option {
	return optionFunc(func(o *options) { o.query["RankDiscount"] = rankdiscount })
}

// WithIDreceiver IDReceiver获取
func (obj *_SalereceiptMgr) WithIDreceiver(idreceiver int) Option {
	return optionFunc(func(o *options) { o.query["IDReceiver"] = idreceiver })
}

// GetByOption 功能选项模式获取
func (obj *_SalereceiptMgr) GetByOption(opts ...Option) (result Salereceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Address").Where("ID = ?", result.IDreceiver).Find(&result.Address).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SalereceiptMgr) GetByOptions(opts ...Option) (results []*Salereceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_SalereceiptMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Salereceipt, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where(options.query)
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
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
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
func (obj *_SalereceiptMgr) GetFromID(id int) (result Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Address").Where("ID = ?", result.IDreceiver).Find(&result.Address).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_SalereceiptMgr) GetBatchFromID(ids []int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateat 通过CreateAt获取内容
func (obj *_SalereceiptMgr) GetFromCreateat(createat int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`CreateAt` = ?", createat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateat 批量查找
func (obj *_SalereceiptMgr) GetBatchFromCreateat(createats []int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`CreateAt` IN (?)", createats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeliverydate 通过DeliveryDate获取内容
func (obj *_SalereceiptMgr) GetFromDeliverydate(deliverydate *int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`DeliveryDate` = ?", deliverydate).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeliverydate 批量查找
func (obj *_SalereceiptMgr) GetBatchFromDeliverydate(deliverydates []*int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`DeliveryDate` IN (?)", deliverydates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIspayment 通过IsPayment获取内容
func (obj *_SalereceiptMgr) GetFromIspayment(ispayment bool) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IsPayment` = ?", ispayment).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIspayment 批量查找
func (obj *_SalereceiptMgr) GetBatchFromIspayment(ispayments []bool) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IsPayment` IN (?)", ispayments).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNote 通过Note获取内容
func (obj *_SalereceiptMgr) GetFromNote(note *string) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`Note` = ?", note).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNote 批量查找
func (obj *_SalereceiptMgr) GetBatchFromNote(notes []*string) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`Note` IN (?)", notes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTypepayment 通过TypePayment获取内容
func (obj *_SalereceiptMgr) GetFromTypepayment(typepayment *string) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`TypePayment` = ?", typepayment).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTypepayment 批量查找
func (obj *_SalereceiptMgr) GetBatchFromTypepayment(typepayments []*string) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`TypePayment` IN (?)", typepayments).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过Status获取内容
func (obj *_SalereceiptMgr) GetFromStatus(status int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`Status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_SalereceiptMgr) GetBatchFromStatus(statuss []int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`Status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTotal 通过Total获取内容
func (obj *_SalereceiptMgr) GetFromTotal(total float32) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`Total` = ?", total).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTotal 批量查找
func (obj *_SalereceiptMgr) GetBatchFromTotal(totals []float32) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`Total` IN (?)", totals).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDcustomer 通过IDCustomer获取内容
func (obj *_SalereceiptMgr) GetFromIDcustomer(idcustomer int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDcustomer 批量查找
func (obj *_SalereceiptMgr) GetBatchFromIDcustomer(idcustomers []int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDCustomer` IN (?)", idcustomers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDdeliver 通过IDDeliver获取内容
func (obj *_SalereceiptMgr) GetFromIDdeliver(iddeliver int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDDeliver` = ?", iddeliver).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDdeliver 批量查找
func (obj *_SalereceiptMgr) GetBatchFromIDdeliver(iddelivers []int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDDeliver` IN (?)", iddelivers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRankdiscount 通过RankDiscount获取内容
func (obj *_SalereceiptMgr) GetFromRankdiscount(rankdiscount float32) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`RankDiscount` = ?", rankdiscount).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRankdiscount 批量查找
func (obj *_SalereceiptMgr) GetBatchFromRankdiscount(rankdiscounts []float32) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`RankDiscount` IN (?)", rankdiscounts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDreceiver 通过IDReceiver获取内容
func (obj *_SalereceiptMgr) GetFromIDreceiver(idreceiver int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDReceiver` = ?", idreceiver).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDreceiver 批量查找
func (obj *_SalereceiptMgr) GetBatchFromIDreceiver(idreceivers []int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDReceiver` IN (?)", idreceivers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
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
func (obj *_SalereceiptMgr) FetchByPrimaryKey(id int) (result Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Address").Where("ID = ?", result.IDreceiver).Find(&result.Address).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexBySalereceiptFk  获取多个内容
func (obj *_SalereceiptMgr) FetchIndexBySalereceiptFk(idcustomer int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexBySalereceiptFk1  获取多个内容
func (obj *_SalereceiptMgr) FetchIndexBySalereceiptFk1(idreceiver int) (results []*Salereceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Salereceipt{}).Where("`IDReceiver` = ?", idreceiver).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Address").Where("ID = ?", results[i].IDreceiver).Find(&results[i].Address).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
