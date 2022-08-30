package LoCred

type LoginCredentialsStruct struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
