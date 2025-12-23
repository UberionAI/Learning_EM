package errorhandling

import (
	"errors"
	"testing"
)

func TestCheckBalance(t *testing.T) {
	tests := []struct {
		name        string
		balance     int
		wantErr     bool
		wantErrType error
	}{
		{
			name:        "+",
			balance:     1000,
			wantErr:     false,
			wantErrType: nil,
		},
		{
			name:        "null",
			balance:     0,
			wantErr:     true,
			wantErrType: ErrInsufficientFunds,
		},
		{
			name:        "-",
			balance:     -50,
			wantErr:     true,
			wantErrType: ErrInsufficientFunds,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := CheckBalance(tt.balance)

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckBalance(%d) ошибка = %v; ожидалась ошибка %t",
					tt.balance, err, tt.wantErr)
				return
			}

			if tt.wantErrType != nil && !errors.Is(err, tt.wantErrType) {
				t.Errorf("CheckBalance(%d) ожидалась ошибка %v, получена %v",
					tt.balance, tt.wantErrType, err)
			}
		})
	}
}