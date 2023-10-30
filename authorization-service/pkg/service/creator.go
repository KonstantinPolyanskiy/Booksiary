package service

import (
	"Booksiary/authorization-service/internal/lib/random"
	"Booksiary/authorization-service/internal/types"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var resChan = make(chan Result)

type Result struct {
	Message string
	Error   error
}

var mu sync.Mutex

var codeDataMap = map[types.Code]types.SavingUser{}

type CreatorService struct {
}

func NewCreatorService() *CreatorService {
	return &CreatorService{}
}

func (cs *CreatorService) UserCode(user types.SavingUser) error {
	user.Code = types.MailAccessCodeData{
		AccessCode: random.Code(),
		ExpiredAt:  time.Now().Add(2 * time.Minute),
	}

	code := types.Code{Code: user.Code.AccessCode}

	log.Print("Код - ", code)

	go AddToMap(code, user)
	go checkCodeExpiration(user)

	return CheckErrorChannel(resChan)
}

func (cs *CreatorService) CheckCode(code types.Code) (types.SavingUser, error) {
	var user types.SavingUser

	if _, ok := codeDataMap[code]; !ok {
		resChan <- Result{
			Message: fmt.Sprintf("Код %v не найден", code),
			Error:   errors.New("провалено"),
		}
		return user, CheckErrorChannel(resChan)
	}
	user = codeDataMap[code]

	return user, nil
}

func checkCodeExpiration(user types.SavingUser) {
	t := time.NewTimer(time.Until(user.Code.ExpiredAt))
	<-t.C

	code := types.Code{Code: user.Code.AccessCode}

	mu.Lock()
	delete(codeDataMap, code)
	mu.Unlock()

	resChan <- Result{
		Message: fmt.Sprintf("Время для кода %v истекло %v назад", user.Code.AccessCode, time.Now().Sub(user.Code.ExpiredAt).Milliseconds()),
		Error:   errors.New("expired"),
	}
}

func AddToMap(c types.Code, user types.SavingUser) {
	mu.Lock()
	codeDataMap[c] = user
	mu.Unlock()
}
