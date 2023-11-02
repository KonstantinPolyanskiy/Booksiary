package service

type Service struct {
	*AuthService
}

func NewService() *Service {
	return &Service{AuthService: NewAuthService()}
}
