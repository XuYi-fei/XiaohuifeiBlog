package blogsDto

type BlogByIdDto struct {
	BlogId string `json:"blogId"`
}

type BlogByUserIdDto struct {
	UserId string `json:"userId"`
}

type ArticleDto struct {
	BlogId         string `json:"blogId"`
	Title          string `json:"title"`
	Digest         string `json:"digest"`
	Content        string `json:"content"`
	UserId         string `json:"userId"`
	Field          string `json:"field"`
	Classification string `json:"classification"`
	Tag            string `json:"tag"`
}

type ArticlesDto struct {
	Body ArticleDto `json:"body"`
}

type BlogUpdateDto struct {
	Id             int    `json:"id,omitempty"`
	BlogId         string `json:"blogId,omitempty"`
	UserId         string `json:"userId,omitempty"`
	Title          string `json:"title,omitempty"`
	Digest         string `json:"digest,omitempty"`
	Content        string `json:"content,omitempty"`
	Author         string `json:"author,omitempty"`
	Field          string `json:"field,omitempty"`
	Classification string `json:"classification,omitempty"`
	Tag            string `json:"tag,omitempty"`
}

// BlogsAllDto 获取所有博客需要的当前页数用来返回博客
type BlogsAllDto struct {
	OrderType      string `json:"orderType"`
	PageNo         int    `json:"pageNo"`
	Mode           string `json:"mode"`
	UserId         string `json:"userId"`
	Classification string `json:"classification"`
	Field          string `json:"field"`
	Tag            string `json:"tag"`
}
