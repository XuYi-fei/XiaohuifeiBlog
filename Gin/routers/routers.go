package routers

import (
	"Gin/controller/adminController"
	"Gin/controller/blogController"
	"Gin/controller/classificationController"
	"Gin/controller/fileController"
	"Gin/controller/sentenceController"
	"Gin/controller/users"
	"Gin/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.Use(middleware.Cors())

	userGroup := r.Group("/users")
	{
		userGroup.POST("/login", users.UserLogin)
		userGroup.POST("/index", middleware.JWTAuth(), users.LoginStatus)
		userGroup.POST("/register", users.UserRegister)
		//userGroup.POST("/getById", middleware.JWTAuth(), users.LoginStatus)
	}

	blogGroup := r.Group("/blogs")
	{
		// 管理员与非管理员上传博客
		blogGroup.POST("/upload", blogController.BlogUpload)
		// 获得最近的博客
		blogGroup.POST("/latest/all", blogController.GetLatestBlogs)
		// 根据BlogId获得博客内容
		blogGroup.POST("/blogId", blogController.GetBlogByBlogId)
		blogGroup.POST("/update", blogController.UpdateBlogById)
		blogGroup.POST("/deleteBlog", blogController.DeleteBlogByTitle)
		// 每页十条数据进行分页返回分页数量
		blogGroup.POST("/getPageNum", blogController.GetBlogsPageNum)
		// 获得展示再首页的各个统计数量
		blogGroup.GET("/getBlogsInfo", blogController.GetBlogsInfo)
	}

	adminGroup := r.Group("/admin")
	{
		// 后台删除一条美句
		adminGroup.POST("/sentences/delete", adminController.DeleteSentencesAdminController)
		// 后台删除一条暂存美句
		adminGroup.POST("/sentencesTmp/delete", adminController.DeleteSentencesTmpAdminController)
		// 后台通过审核通过一条美句
		adminGroup.POST("/sentencesTmp/pass", adminController.PassSentencesTmpAdminController)
		// 后台删除一条暂存博客
		adminGroup.POST("/blogsTmp/delete", adminController.DeleteBlogsTmpAdminController)
		// 后台通过一条博客
		adminGroup.POST("/blogsTmp/pass", adminController.PassBlogsTmpAdminController)
	}

	sentenceGroup := r.Group("/sentences")
	{
		// 根据总数量按照每页十条数据进行分页返回分页数量
		sentenceGroup.POST("/getPageNum", sentenceController.GetSentencesPageNum)
		// 上传一条数据，包括管理员和非管理员
		sentenceGroup.POST("/upload", sentenceController.SentenceUpload)
		// 获取所有的句子
		sentenceGroup.POST("/all", sentenceController.GetAllSentences)
		// 获取所有的轮播句子
		sentenceGroup.GET("/show", sentenceController.GetSentencesShow)
		// 用户管理自己的美句进行删除
		sentenceGroup.POST("/delete/Personal", sentenceController.DeleteSentencesPersonalController)
	}

	classificationGroup := r.Group("/classification")
	{
		classificationGroup.POST("/getField", classificationController.GetField)
		classificationGroup.POST("/getClassification", classificationController.GetClassification)
		classificationGroup.POST("/getTagBlogs", classificationController.GetTagBlogs)
	}

	filesGroup := r.Group("/files")
	{
		filesGroup.POST("/images", fileController.ImagesUpload)
	}
	return r
}
