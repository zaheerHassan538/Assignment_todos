package request

type CreateUsersRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
