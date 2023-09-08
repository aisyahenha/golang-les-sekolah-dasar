package mappingutil

import "github.com/aisyahenha/golang-les-sekolah-dasar/model"

// type MappingUser interface{
// 	MapingUserResponse (user *model.User)
// }

func MappingUser(user model.User) ( model.UserRespon) {
	var userRes model.UserRespon = model.UserRespon{
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,	
		user.Username,
		user.Role,
		user.ResetToken,		
	}
	
	return userRes
}

/*
type UserRespon struct {
	BaseModel
	Username   string `json:"username"`
	Role       string `json:"role"`
	ResetToken string `json:"resetToken,omitempty"`
}

*/
