package time_day

import (
	"testing"
	"time"
)

func TestIsWorkMorning(t *testing.T) {
	tests := []struct {
		name     string
		fixedNow time.Time
		want     bool
	}{
		{
			name:     "понедельник 10:00",
			fixedNow: time.Date(2025, 12, 22, 10, 0, 0, 0, time.UTC),
			want:     true,
		},
		{
			name:     "понедельник 14:00",
			fixedNow: time.Date(2025, 12, 22, 14, 0, 0, 0, time.UTC),
			want:     false,
		},
		{
			name:     "суббота 10:00",
			fixedNow: time.Date(2025, 12, 27, 10, 0, 0, 0, time.UTC),
			want:     false,
		},
		{
			name:     "воскресенье 10:00",
			fixedNow: time.Date(2025, 12, 28, 10, 0, 0, 0, time.UTC),
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalNow := now
			now = func() time.Time { return tt.fixedNow }
			defer func() { now = originalNow }()

			if got := IsWorkMorning(); got != tt.want {
				t.Errorf("IsWorkMorning() = %v, want %v", got, tt.want)
			}
		})
	}
}
