package model

type CourseModel struct {
	BaseModel
	Subject string `json:"subject"`
	Class   string `json:"class"`
	Day     string `json:"day"`
	Start   string `json:"start"`
	End     string `json:"end"`
}
