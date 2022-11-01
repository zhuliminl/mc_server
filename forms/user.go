package forms

type UserCreate struct {
	Username string `json:"username" binding:"required,max=10"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDelete struct {
	UserId string `json:"userId" binding:"required"`
}
