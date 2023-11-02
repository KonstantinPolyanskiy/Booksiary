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

type SavingUserMap struct {
	mu      *sync.Mutex
	UserMap map[types.Code]types.SavingUser
}

func (sm *SavingUserMap) Add(code types.Code, user types.SavingUser) {
	sm.mu.Lock()
	sm.UserMap[code] = user
	sm.mu.Unlock()
}

type Adder interface {
	Add(code types.Code, user types.SavingUser)
}

type Result struct {
	Message string
	Error   error
}

var mu sync.Mutex

var codeDataMap = map[types.Code]types.SavingUser{}

type CreatorService struct {
	userMap SavingUserMap
}

func NewCreatorService(um SavingUserMap) *CreatorService {
	return &CreatorService{
		userMap: um,
	}
}

func (cs *CreatorService) UserCode(user types.UserCreateResponse) error {
	savingUser := types.SavingUser{
		User: types.User{
			Uuid: "",
			Personality: types.Personality{
				Name:    user.Name,
				Surname: user.Surname,
			},
		},
		Code: types.MailAccessCodeData{},
	}
	savingUser.Code = types.MailAccessCodeData{
		AccessCode: random.Code(),
		ExpiredAt:  time.Now().Add(2 * time.Minute),
	}

	code := types.Code{Code: savingUser.Code.AccessCode}

	log.Print("Код - ", code)

	cs.userMap.Add(code, savingUser)
	CheckCodeExpiration(savingUser, cs.userMap.UserMap)

	return CheckErrorChannel(resChan)
}

func (cs *CreatorService) FindUserByCode(code types.Code) (types.SavingUser, error) {
	user, err := searchUser(code, cs.userMap.UserMap)
	if err != nil {
		return types.SavingUser{}, err
	}

	return user, nil
}

func AddToMap(c types.Code, user types.SavingUser) {
	mu.Lock()
	codeDataMap[c] = user
	mu.Unlock()
}

// searchUser ищет по переданному ключу пользователя в карте
func searchUser(code types.Code, m map[types.Code]types.SavingUser) (types.SavingUser, error) {
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
