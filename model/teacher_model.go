package model

type TeacherModel struct {
	BaseModel
	BasePersonalModel
	Specialist string `json:"specialist"`
}
