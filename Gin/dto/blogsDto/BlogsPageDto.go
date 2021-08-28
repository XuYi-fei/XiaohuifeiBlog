package blogsDto

type PersonalBlogsPageDto struct {
	UserId string `json:"userId"`
}

type BlogPageDto struct {
	UserId string `json:"userId"`
	Mode   string `json:"mode"`
}
