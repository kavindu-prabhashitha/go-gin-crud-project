package services

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUserName string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUserName: "kavindu",
		authorizedPassword: "kp@123",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUserName == username && service.authorizedPassword == password
}
