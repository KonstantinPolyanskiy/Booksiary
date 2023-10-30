package service

import "log"

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
