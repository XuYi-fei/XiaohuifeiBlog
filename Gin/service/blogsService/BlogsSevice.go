package blogsService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/blogsDto"
	"Gin/models"
	"Gin/vo/blogsVO"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

const (
	//BLOGS_BASE = "D:/GitRepos/blog-go/Gin/resources/blogs"
	BLOGS_BASE = "../../resources/blogs"
)

func BlogsUploadService(dto *blogsDto.ArticlesDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := dto.Body.UserId
	var user models.Users
	var oldBlog models.Article
	var tmpOldBlog models.ArticleTmp
	newBlog := new(models.Article)
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.LoginStatusInvalid, returnMsg.LoginStatusInvalid, errors.New("用户登录失效"), nil
	}

	dao.DB.Where("title = ?", dto.Body.Title).Find(&oldBlog)
	dao.DB.Where("title = ?", dto.Body.Title).Find(&tmpOldBlog)
	if oldBlog.Title != "" || tmpOldBlog.Title != "" {
		return returnCode.TitleInvalid, returnMsg.BlogTitleDuplicate, errors.New("标题已经存在,请更新标题后重试\""), nil
	}

	// 根据title生成blog的blog-id
	hash := sha256.New()
	hash.Write([]byte(dto.Body.Title))
	blogId := hex.EncodeToString(hash.Sum(nil))
	// 存储文章内容部分，将其存在文件中
	filePath := path.Join(BLOGS_BASE, UserId)
	fileName := path.Join(filePath, blogId+".txt")
	// 检测文件是否已经创建，如果不存在就自动创建
	if _, err := os.Stat(filePath); err != nil {
		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			return returnCode.BlogCreateFailed, returnMsg.BlogCreateFailed, err, nil
		}
	}
	// 写入文件
	content := dto.Body.Content
	if ioutil.WriteFile(fileName, []byte(content), 0777) != nil {
		return returnCode.BlogCreateFailed, returnMsg.BlogCreateFailed, errors.New("博客文章更新失败，请再试一次"), nil
	}
	// 生成新文章
	{
		newBlog.UserId = dto.Body.UserId
		newBlog.BlogId = blogId
		newBlog.Author = user.UserName
		newBlog.Title = dto.Body.Title
		newBlog.Content = fileName
		newBlog.Digest = dto.Body.Digest
		newBlog.Field = dto.Body.Field
		newBlog.Classification = dto.Body.Classification
		newBlog.Tag = dto.Body.Tag
	}
	if user.Privilege == 1 {
		dao.DB.Create(&newBlog)
		user.BlogNum++
		dao.DB.Updates(&user)
		return returnCode.OK, returnMsg.BlogUploadSuccess, nil, nil
	} else {
		tmpBlog := models.NewArticleTmp(newBlog.BlogId, newBlog.UserId, newBlog.Title, newBlog.Digest, newBlog.Content,
			newBlog.Author, newBlog.Field, newBlog.Classification, newBlog.Tag)
		dao.DB.Create(&tmpBlog)
		return returnCode.OK, returnMsg.BlogUploadWaitingGroup, nil, nil
	}
}

func GetLatestBlogsService(blogDto *blogsDto.BlogsAllDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	//blogs := reflect.New(reflect.TypeOf(models.NewBlogs()))
	//var blogs interface{}
	//switch blogDto.Mode {
	//case constant.AllBlogs:
	//	var s = blogs.(models.Article)
	//	fmt.Println(s)
	//	blogs = blogs.(models.Article)
	//case constant.PersonalBlogs:
	//	blogs = blogs.(*[]models.Article)
	//case constant.AllBlogsTmp:
	//	blogs = blogs.(*[]models.ArticleTmp)
	//}
	//fmt.Println(blogs)
	if blogDto.Mode == constant.AllBlogs {
		var blogs []models.Article
		dao.DB.Order("updated_at desc").Offset((blogDto.PageNo - 1) * 10).Limit((blogDto.PageNo) * 10).Find(&blogs)
		if len(blogs) == 0 {
			return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), nil
		}
		return returnCode.OK, returnMsg.OK, nil, blogs
	} else if blogDto.Mode == constant.PersonalBlogs {
		var blogs []models.Article
		dao.DB.Where("user_id = ?", blogDto.UserId).Order("updated_at desc").Offset((blogDto.PageNo - 1) * 10).Limit((blogDto.PageNo) * 10).Find(&blogs)
		if len(blogs) == 0 {
			return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), nil
		}
		return returnCode.OK, returnMsg.OK, nil, blogs
	} else if blogDto.Mode == constant.AllBlogsTmp {
		var blogs []models.ArticleTmp
		dao.DB.Order("id desc").Offset((blogDto.PageNo - 1) * 10).Limit((blogDto.PageNo) * 10).Find(&blogs)
		if len(blogs) == 0 {
			return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), nil
		}
		for i := 0; i < len(blogs); i++ {
			fileName := blogs[i].Content
			buf, err := ioutil.ReadFile(fileName)
			if err != nil {
				//dao.DB.Where("id = ?", blogs[i].ID).Delete(&models.Article{})
			}
			blogs[i].Content = string(buf)
		}
		return returnCode.OK, returnMsg.OK, nil, blogs
	}
	return returnCode.ModeInvalid, returnMsg.ModeInvalid, errors.New("mode无效"), nil

}

func GetBlogByBlogIDService(blogId string) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	blogVo := new(blogsVO.BlogVO)
	var blog models.Article
	dao.DB.Where("blog_id = ?", blogId).Find(&blog)
	if blog.BlogId == "" {
		return returnCode.BlogIdInvalid, returnMsg.BlogIdInvalid, errors.New("文章不存在"), nil
	}

	{
		fileName := blog.Content
		buf, err := ioutil.ReadFile(fileName)
		if err != nil {
			dao.DB.Where("id = ?", blog.Id).Delete(&models.Article{})
			return returnCode.BlogIdInvalid, returnMsg.BlogIdInvalid, errors.New("文章对应txt未找到"), nil
		}
		blogVo = blogsVO.NewBlogVOFromModel(blog, string(buf))
	}
	return returnCode.OK, returnMsg.BlogFoundSuccess, nil, blogVo
}

func UpdateBlogByIdService(dto *blogsDto.ArticleDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := dto.UserId
	var user models.Users
	var oldBlog models.Article
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.LoginStatusInvalid, returnMsg.LoginStatusInvalid, errors.New("登录失效"), nil
	}

	dao.DB.Where("title = ?", dto.Title).Find(&oldBlog)

	// 根据title生成blog的blog-id
	hash := sha256.New()
	hash.Write([]byte(dto.Title))
	blogId := hex.EncodeToString(hash.Sum(nil))
	// 存储文章内容部分，将其存在文件中
	filePath := path.Join(BLOGS_BASE, UserId)
	fileName := path.Join(filePath, blogId+".txt")
	// 检测文件是否已经创建，如果不存在就自动创建
	if _, err := os.Stat(filePath); err != nil {
		err := os.MkdirAll(filePath, 0777)
		return returnCode.BlogUpdateFailed, returnMsg.BlogUpdateFailed, err, nil
	}
	// 写入文件
	content := dto.Content
	if ioutil.WriteFile(fileName, []byte(content), 0777) != nil {
		return returnCode.BlogUpdateFailed, returnMsg.BlogUpdateFailed, errors.New("写入文件失败"), nil
	}
	// 生成新文章
	{
		oldBlog.UserId = dto.UserId
		oldBlog.BlogId = blogId
		oldBlog.Author = user.UserName
		oldBlog.Title = dto.Title
		oldBlog.Content = fileName
		oldBlog.Digest = dto.Digest
		oldBlog.Field = dto.Field
		oldBlog.Classification = dto.Classification
		oldBlog.Tag = dto.Tag
	}

	dao.DB.Updates(&oldBlog)
	return returnCode.BlogUpdateSuccess, returnMsg.BlogUpdateSuccess, nil, nil
}

func DeleteBlogByTitleService(dto *blogsDto.ArticleDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := dto.UserId
	var user models.Users
	var oldBlog models.Article
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.LoginStatusInvalid, returnMsg.LoginStatusInvalid, errors.New("登录失效"), nil
	}

	filePath := path.Join(BLOGS_BASE, UserId)
	fileName := path.Join(filePath, dto.BlogId+".txt")
	err := os.Remove(fileName)
	if err != nil {
		return returnCode.BlogDeleteFailed, returnMsg.BlogDeleteFailed, err, nil
	}
	dao.DB.Where("title = ?", dto.Title).Delete(&oldBlog)
	return returnCode.OK, returnMsg.BlogDeleteSuccess, nil, nil
}

func PassBlogsService(blogDto *blogsDto.BlogUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := blogDto.UserId
	var user models.Users
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.UserNotFound, returnMsg.UserNotFound, errors.New("用户不存在"), nil
	}

	newBlog := models.NewArticle(blogDto.BlogId, blogDto.UserId, blogDto.Title, blogDto.Digest, blogDto.Content,
		blogDto.Author, blogDto.Field, blogDto.Classification, blogDto.Tag)
	dao.DB.Create(&newBlog)
	dao.DB.Where("id = ?", blogDto.Id).Delete(&models.ArticleTmp{})
	return returnCode.OK, returnMsg.CheckPassed, nil, nil
}

func DeleteBlogsService(blogDto *blogsDto.BlogUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	dao.DB.Where("id = ?", blogDto.Id).Delete(&models.ArticleTmp{})
	return returnCode.OK, returnMsg.BlogDeleteSuccess, nil, nil
}

func GetBlogsInfoService() (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var result blogsVO.BlogInfos
	dao.DB.Raw("select count(*) as BlogNum, count(distinct field) as FieldNum, count(distinct classification) as ClassificationNum, " +
		"count(distinct tag) as TagNum from articles").Find(&result)
	return returnCode.OK, returnMsg.OK, nil, result
}
