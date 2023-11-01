package service

type Service struct {
	Creator
	Updater
	Deleter
	Getter
}

func NewService() *Service {
	return &Service{
		Creator: NewCreator(),
		Updater: NewUpdateService(),
		Deleter: NewDeleteService(),
		Getter:  NewGetterService(),
	}
}
