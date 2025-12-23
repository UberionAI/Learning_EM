package concurrency

import (
	"sync"
	"testing"
)

func TestCounterRaceCondition(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()

	got := counter.Value()
	t.Logf("Non-atomic = %d (ожидаем 1000, но race!)", got)

	if got != 1000 {
		t.Errorf("FAIL: %d != 1000 (race detected!)", got)
	}
}

func TestCounterAtomic(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.AddAtomic(1)
		}()
	}
	wg.Wait()

	if counter.Value() != 1000 {
		t.Errorf("Atomic FAIL: %d != 1000", counter.Value())
	}
}
