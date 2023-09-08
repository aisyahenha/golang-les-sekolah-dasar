package model

type ScheduleModel struct {
	BaseModel
	CourseID  string `json:"courseid"`
	Course    CourseModel `gorm:"foreignKey:CourseID"`
	TeacherID string `json:"teacherid"`
	Teacher   TeacherModel `gorm:"foreignKey:TeacherID"`
	StudentID string `json:"studentid"`
	Student   StudentModel `gorm:"foreignKey:StudentID"`
}
