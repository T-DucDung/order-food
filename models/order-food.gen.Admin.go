package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AdminMgr struct {
	*_BaseMgr
}

// AdminMgr open func
func AdminMgr(db *gorm.DB) *_AdminMgr {
	if db == nil {
		panic(fmt.Errorf("AdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Admin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminMgr) GetTableName() string {
	return "Admin"
}

// Reset 重置gorm会话
func (obj *_AdminMgr) Reset() *_AdminMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AdminMgr) Get() (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AdminMgr) Gets() (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Admin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_AdminMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_AdminMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithUsername UserName获取
func (obj *_AdminMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["UserName"] = username })
}

// WithPass Pass获取
func (obj *_AdminMgr) WithPass(pass string) Option {
	return optionFunc(func(o *options) { o.query["Pass"] = pass })
}

// WithPhone Phone获取
func (obj *_AdminMgr) WithPhone(phone *int) Option {
	return optionFunc(func(o *options) { o.query["Phone"] = phone })
}

// WithAddress Address获取
func (obj *_AdminMgr) WithAddress(address *string) Option {
	return optionFunc(func(o *options) { o.query["Address"] = address })
}

// WithEmail Email获取
func (obj *_AdminMgr) WithEmail(email *string) Option {
	return optionFunc(func(o *options) { o.query["Email"] = email })
}

// WithSex Sex获取
func (obj *_AdminMgr) WithSex(sex *bool) Option {
	return optionFunc(func(o *options) { o.query["Sex"] = sex })
}

// WithStatus Status获取
func (obj *_AdminMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["Status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AdminMgr) GetByOption(opts ...Option) (result Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminMgr) GetByOptions(opts ...Option) (results []*Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AdminMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Admin, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query)
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
func (obj *_AdminMgr) GetFromID(id int) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`ID` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AdminMgr) GetBatchFromID(ids []int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`ID` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过Name获取内容
func (obj *_AdminMgr) GetFromName(name string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AdminMgr) GetBatchFromName(names []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Name` IN (?)", names).Find(&results).Error

	return
}

// GetFromUsername 通过UserName获取内容
func (obj *_AdminMgr) GetFromUsername(username string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`UserName` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找
func (obj *_AdminMgr) GetBatchFromUsername(usernames []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`UserName` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPass 通过Pass获取内容
func (obj *_AdminMgr) GetFromPass(pass string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Pass` = ?", pass).Find(&results).Error

	return
}

// GetBatchFromPass 批量查找
func (obj *_AdminMgr) GetBatchFromPass(passs []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Pass` IN (?)", passs).Find(&results).Error

	return
}

// GetFromPhone 通过Phone获取内容
func (obj *_AdminMgr) GetFromPhone(phone *int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_AdminMgr) GetBatchFromPhone(phones []*int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddress 通过Address获取内容
func (obj *_AdminMgr) GetFromAddress(address *string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找
func (obj *_AdminMgr) GetBatchFromAddress(addresss []*string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromEmail 通过Email获取内容
func (obj *_AdminMgr) GetFromEmail(email *string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_AdminMgr) GetBatchFromEmail(emails []*string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromSex 通过Sex获取内容
func (obj *_AdminMgr) GetFromSex(sex *bool) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找
func (obj *_AdminMgr) GetBatchFromSex(sexs []*bool) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromStatus 通过Status获取内容
func (obj *_AdminMgr) GetFromStatus(status bool) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AdminMgr) GetBatchFromStatus(statuss []bool) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`Status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AdminMgr) FetchByPrimaryKey(id int) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`ID` = ?", id).First(&result).Error

	return
}
