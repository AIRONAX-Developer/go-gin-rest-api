package service

type LoginService interface {
	LoginUser(email, password string) bool
}

type LoginInformationStruct struct {
	email    string
	password string
}

func StaticLoginService() LoginService {

	return &LoginInformationStruct{
		email:    "info@chinmayvivek.com",
		password: "test",
	}
}
func (loginInfo *LoginInformationStruct) LoginUser(email, password string) bool {
	return loginInfo.email == email && loginInfo.password == password
}
