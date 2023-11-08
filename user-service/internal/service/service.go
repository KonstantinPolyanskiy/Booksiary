package service

type Service struct {
	*SignUpService
	Updater
	Deleter
	Getter
}

func NewService() *Service {
	return &Service{
		SignUpService: NewSignUpService(),
		Updater:       NewUpdateService(),
		Deleter:       NewDeleteService(),
		Getter:        NewGetterService(),
	}
}
