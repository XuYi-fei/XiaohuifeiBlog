package returnMsg

type ReturnMsg string

const (
	OK                     ReturnMsg = "操作成功"
	JsonInvalid                      = "Json数据解析失败"
	LoginStatusInvalid               = "登录状态失效"
	UploadSuccessful                 = "上传成功"
	UploadWaiting                    = "提交成功，等待后台审核"
	UserIdDuplicate                  = "用户ID已存在"
	UserNameDuplicate                = "用户名已存在"
	BlogTitleDuplicate               = "文章标题已存在"
	BlogIdInvalid                    = "文章不存在"
	BlogUpdateFailed                 = "博客更新失败"
	BlogUpdateSuccess                = "博客更新成功"
	BlogUploadSuccess                = "文章发布成功"
	BlogUploadWaitingGroup           = "文章已提交后台，等待审核"
	BlogCreateFailed                 = "文章创建失败"
	BlogCreateSuccess                = "文章创建成功"
	BlogDeleteFailed                 = "文章删除失败"
	BlogDeleteSuccess                = "文章删除成功"
	BlogFoundSuccess                 = "查找文章成功"
	SentenceDeleteSuccess            = "文章删除成功"
	UserNotFound                     = "用户不存在"
	CheckPassed                      = "审核通过"
	RegisterOK                       = "注册成功"
	NoData                           = "暂无数据"
	ModeInvalid                      = "Mode无效"
)
