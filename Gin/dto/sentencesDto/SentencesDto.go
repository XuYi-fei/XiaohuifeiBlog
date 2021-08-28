package sentencesDto

type SentenceUploadDto struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type SentencesAllPersonalDto struct {
	PageNo int    `json:"pageNo"`
	UserId string `json:"userId"`
}

type SentenceUpdateDto struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

// BlogsAllDto 获取所有博客需要的当前页数用来返回博客
type SentencesAllDto struct {
	OrderType string `json:"orderType"`
	PageNo    int    `json:"pageNo"`
	Mode      string `json:"mode"`
	UserId    string `json:"userId"`
}
