package model
//true - Male
//false - Female
type BasePersonalModel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  bool   `json:"gender"`
}
