package service

import (
	"Booksiary/authorization-service/internal/types"
	"errors"
	"fmt"
	"log"
	"time"
)

func CheckErrorChannel(ch <-chan Result) error {
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				return nil
			}
			if r.Error != nil {
				log.Print(r.Message)
				return r.Error
			}
		default:
			return nil
		}
	}
}

// CheckCodeExpiration проверяет пользователей в карте и записывает их в канал
func CheckCodeExpiration(user types.SavingUser, m map[types.Code]types.SavingUser) {
	go func() {
		t := time.NewTimer(time.Until(user.Code.ExpiredAt))
		<-t.C

		code := types.Code{Code: user.Code.AccessCode}

		resChan <- Result{
			Message: fmt.Sprintf("Время для кода %v истекло %v назад", user.Code.AccessCode, time.Now().Sub(user.Code.ExpiredAt).Milliseconds()),
			Error:   errors.New("expired"),
		}

		mu.Lock()
		delete(m, code)
		mu.Unlock()
	}()
}
