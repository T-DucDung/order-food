package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CommentMgr struct {
	*_BaseMgr
}

// CommentMgr open func
func CommentMgr(db *gorm.DB) *_CommentMgr {
	if db == nil {
		panic(fmt.Errorf("CommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CommentMgr) GetTableName() string {
	return "Comment"
}

// Reset 重置gorm会话
func (obj *_CommentMgr) Reset() *_CommentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CommentMgr) Get() (result Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_CommentMgr) Gets() (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CommentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Comment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID ID获取
func (obj *_CommentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["ID"] = id })
}

// WithContent Content获取
func (obj *_CommentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["Content"] = content })
}

// WithCreateat CreateAt获取
func (obj *_CommentMgr) WithCreateat(createat int) Option {
	return optionFunc(func(o *options) { o.query["CreateAt"] = createat })
}

// WithIDproduct IDProduct获取
func (obj *_CommentMgr) WithIDproduct(idproduct int) Option {
	return optionFunc(func(o *options) { o.query["IDProduct"] = idproduct })
}

// WithIDcustomer IDCustomer获取
func (obj *_CommentMgr) WithIDcustomer(idcustomer int) Option {
	return optionFunc(func(o *options) { o.query["IDCustomer"] = idcustomer })
}

// WithRate Rate获取
func (obj *_CommentMgr) WithRate(rate int) Option {
	return optionFunc(func(o *options) { o.query["Rate"] = rate })
}

// GetByOption 功能选项模式获取
func (obj *_CommentMgr) GetByOption(opts ...Option) (result Comment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CommentMgr) GetByOptions(opts ...Option) (results []*Comment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_CommentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Comment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Comment{}).Where(options.query)
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
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
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
func (obj *_CommentMgr) GetFromID(id int) (result Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_CommentMgr) GetBatchFromID(ids []int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`ID` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromContent 通过Content获取内容
func (obj *_CommentMgr) GetFromContent(content string) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`Content` = ?", content).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromContent 批量查找
func (obj *_CommentMgr) GetBatchFromContent(contents []string) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`Content` IN (?)", contents).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateat 通过CreateAt获取内容
func (obj *_CommentMgr) GetFromCreateat(createat int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`CreateAt` = ?", createat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateat 批量查找
func (obj *_CommentMgr) GetBatchFromCreateat(createats []int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`CreateAt` IN (?)", createats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDproduct 通过IDProduct获取内容
func (obj *_CommentMgr) GetFromIDproduct(idproduct int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDproduct 批量查找
func (obj *_CommentMgr) GetBatchFromIDproduct(idproducts []int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`IDProduct` IN (?)", idproducts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDcustomer 通过IDCustomer获取内容
func (obj *_CommentMgr) GetFromIDcustomer(idcustomer int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDcustomer 批量查找
func (obj *_CommentMgr) GetBatchFromIDcustomer(idcustomers []int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`IDCustomer` IN (?)", idcustomers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRate 通过Rate获取内容
func (obj *_CommentMgr) GetFromRate(rate int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`Rate` = ?", rate).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRate 批量查找
func (obj *_CommentMgr) GetBatchFromRate(rates []int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`Rate` IN (?)", rates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
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
func (obj *_CommentMgr) FetchByPrimaryKey(id int) (result Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`ID` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Product").Where("ID = ?", result.IDproduct).Find(&result.Product).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("User").Where("ID = ?", result.IDcustomer).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByCommentFk  获取多个内容
func (obj *_CommentMgr) FetchIndexByCommentFk(idproduct int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`IDProduct` = ?", idproduct).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByCommentFk1  获取多个内容
func (obj *_CommentMgr) FetchIndexByCommentFk1(idcustomer int) (results []*Comment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comment{}).Where("`IDCustomer` = ?", idcustomer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("Product").Where("ID = ?", results[i].IDproduct).Find(&results[i].Product).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("User").Where("ID = ?", results[i].IDcustomer).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
