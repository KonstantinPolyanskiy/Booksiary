package service

import (
	"Booksiary/authorization-service/internal/lib/random"
	"Booksiary/authorization-service/internal/types"
	"errors"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

var ResChan = make(chan Result, 10)

type Result struct {
	Message string
	Error   error
}

var expiredChan = make(chan bool)

var mu sync.Mutex

var CodeDataMap = map[types.Code]types.SaveUser{}

type CreatorService struct {
}

func NewCreatorService() *CreatorService {
	return &CreatorService{}
}

func (cs *CreatorService) UserCode(user types.SaveUser) error {
	user.CodeData = types.CodeData{
		Code:      random.Code(),
		ExpiredAt: time.Now().Add(2 * time.Minute),
	}

	log.Print("Код - ", user.Code)

	go AddToMap(types.Code{Code: user.Code}, user)
	go checkCodeExpiration(user)

	return CheckErrorChannel(ResChan)
}

func (cs *CreatorService) CheckCode(c types.Code) (types.SaveUser, error) {
	var user types.SaveUser

	if _, ok := CodeDataMap[c]; !ok {
		ResChan <- Result{
			Message: fmt.Sprintf("Код %v не найден", c),
			Error:   errors.New("провалено"),
		}
		return user, CheckErrorChannel(ResChan)
	}
	user = CodeDataMap[c]

	return user, nil
}

func checkCodeExpiration(user types.SaveUser) {
	t := time.NewTimer(time.Until(user.ExpiredAt))
	<-t.C
	mu.Lock()
	delete(CodeDataMap, types.Code{Code: user.Code})
	mu.Unlock()
	ResChan <- Result{
		Message: fmt.Sprintf("Время для кода %v истекло %v", user.Code, time.Now().Sub(user.ExpiredAt)),
		Error:   errors.New("expired"),
	}
}

func AddToMap(c types.Code, user types.SaveUser) {
	var ms runtime.MemStats
	mu.Lock()
	CodeDataMap[c] = user
	mu.Unlock()

	runtime.ReadMemStats(&ms)

	fmt.Println("Памяти - ", ms.Alloc/1024/1024)
}
