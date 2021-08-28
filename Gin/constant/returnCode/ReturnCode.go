package returnCode

type ReturnCode int

const (
	OK ReturnCode = iota
	BadRequest
	Expired
	LoginStatusInvalid
	TitleInvalid
	NoData
	BlogIdInvalid
	BlogCreateFailed   // 博客创建失败: 创建文件夹失败
	BlogUpdateFailed   // 博客更新失败
	BlogUpdateSuccess  // 博客更新成功
	BlogDeleteFailed   // 博客删除失败
	BlogDeleteSuccess  // 博客删除成功
	ArticleNotFound    // txt文件不存在
	NoPrivilege        // 权限不足
	JsonInvalid        // 传输数据格式有误
	UserIdDuplicate    // 注册时用户ID重复
	UserNameDuplicate  // 注册时用户名重复
	UserNotFound       // 用户不存在
	DeleteError        // 删除博客失败
	ModeInvalid        // 查询时选择的模式不正确，当前可选项: 全部博客、全部待审核博客、某人的全部博客
	UploadImageInvalid //上传图片时图片无效
)
