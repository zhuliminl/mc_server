package forms

type UserSignUp struct {
	Name     string `json:"name" binding:"required,max=10"`
	BirthDay string `json:"birthday" binding:"required"`
}

type UserId struct {
	id string `form:"id" binding:"required"`
}
