package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("User"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "User"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Rank").Where("ID = ?", result.IDrank).Find(&result.Rank).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_UserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithName Name获取
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["Name"] = name })
}

// WithUsername Username获取
func (obj *_UserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["Username"] = username })
}

// WithPass Pass获取
func (obj *_UserMgr) WithPass(pass string) Option {
	return optionFunc(func(o *options) { o.query["Pass"] = pass })
}

// WithPhone Phone获取
func (obj *_UserMgr) WithPhone(phone *string) Option {
	return optionFunc(func(o *options) { o.query["Phone"] = phone })
}

// WithAddress Address获取
func (obj *_UserMgr) WithAddress(address *string) Option {
	return optionFunc(func(o *options) { o.query["Address"] = address })
}

// WithEmail Email获取
func (obj *_UserMgr) WithEmail(email *string) Option {
	return optionFunc(func(o *options) { o.query["Email"] = email })
}

// WithSex Sex获取
func (obj *_UserMgr) WithSex(sex bool) Option {
	return optionFunc(func(o *options) { o.query["Sex"] = sex })
}

// WithCreateat CreateAt获取
func (obj *_UserMgr) WithCreateat(createat int) Option {
	return optionFunc(func(o *options) { o.query["CreateAt"] = createat })
}

// WithStatus Status获取
func (obj *_UserMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["Status"] = status })
}

// WithIDrank IDRank获取
func (obj *_UserMgr) WithIDrank(idrank int) Option {
	return optionFunc(func(o *options) { o.query["IDRank"] = idrank })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Rank").Where("ID = ?", result.IDrank).Find(&result.Rank).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_UserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]User, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
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
func (obj *_UserMgr) GetFromID(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Rank").Where("ID = ?", result.IDrank).Find(&result.Rank).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_UserMgr) GetBatchFromID(ids []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过Name获取内容
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUsername 通过Username获取内容
func (obj *_UserMgr) GetFromUsername(username string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Username` = ?", username).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUsername 批量查找
func (obj *_UserMgr) GetBatchFromUsername(usernames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Username` IN (?)", usernames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPass 通过Pass获取内容
func (obj *_UserMgr) GetFromPass(pass string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Pass` = ?", pass).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPass 批量查找
func (obj *_UserMgr) GetBatchFromPass(passs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Pass` IN (?)", passs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPhone 通过Phone获取内容
func (obj *_UserMgr) GetFromPhone(phone *string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Phone` = ?", phone).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPhone 批量查找
func (obj *_UserMgr) GetBatchFromPhone(phones []*string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Phone` IN (?)", phones).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAddress 通过Address获取内容
func (obj *_UserMgr) GetFromAddress(address *string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Address` = ?", address).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAddress 批量查找
func (obj *_UserMgr) GetBatchFromAddress(addresss []*string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Address` IN (?)", addresss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEmail 通过Email获取内容
func (obj *_UserMgr) GetFromEmail(email *string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Email` = ?", email).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEmail 批量查找
func (obj *_UserMgr) GetBatchFromEmail(emails []*string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Email` IN (?)", emails).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSex 通过Sex获取内容
func (obj *_UserMgr) GetFromSex(sex bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Sex` = ?", sex).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSex 批量查找
func (obj *_UserMgr) GetBatchFromSex(sexs []bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Sex` IN (?)", sexs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateat 通过CreateAt获取内容
func (obj *_UserMgr) GetFromCreateat(createat int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`CreateAt` = ?", createat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateat 批量查找
func (obj *_UserMgr) GetBatchFromCreateat(createats []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`CreateAt` IN (?)", createats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过Status获取内容
func (obj *_UserMgr) GetFromStatus(status bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_UserMgr) GetBatchFromStatus(statuss []bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`Status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDrank 通过IDRank获取内容
func (obj *_UserMgr) GetFromIDrank(idrank int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`IDRank` = ?", idrank).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDrank 批量查找
func (obj *_UserMgr) GetBatchFromIDrank(idranks []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`IDRank` IN (?)", idranks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
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
func (obj *_UserMgr) FetchByPrimaryKey(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Rank").Where("ID = ?", result.IDrank).Find(&result.Rank).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByUserFk  获取多个内容
func (obj *_UserMgr) FetchIndexByUserFk(idrank int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`IDRank` = ?", idrank).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Rank").Where("ID = ?", results[i].IDrank).Find(&results[i].Rank).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
