package errorhandling

import "errors"

var ErrInsufficientFunds = errors.New("недостаточно средств на счете")

func CheckBalance(balance int) error {
	if balance <= 0 {
		return ErrInsufficientFunds
	}
	return nil
}