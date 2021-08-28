package classificationDto

type FieldDto struct {
	UserId string `json:"userId"`
	Mode   string `json:"mode"`
}

type ClassificationDto struct {
	UserId string `json:"userId"`
	Mode   string `json:"mode"`
	Field  string `json:"field"`
}

type TagDto struct {
	ClassificationDto
	Classification string `json:"classification"`
}
