package sentencesDto

type PersonalSentencePageDto struct {
	UserId string `json:"userId"`
}

type SentencePageDto struct {
	UserId string `json:"userId"`
	Mode   string `json:"mode"`
}
