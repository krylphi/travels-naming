package naming

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

// TestOutput these are not actual functional unit tests, just some output testing
func TestOutput(t *testing.T) {
	staticTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t.Run("all flips false", func(t *testing.T) {
		random := newRandomizer(staticTime)
		f := func() bool {
			return false
		}
		random.flipper = &f
		for i := 0; i < 10000; i++ {
			t1, t2 := randTime(), randTime()
			log.Printf("now: %s, event: %s, epithet: %s",
				t1.Format(time.RFC3339), t2.Format(time.RFC3339), epithetFunction(random, t1, t2, "New York"))
		}
	})
	t.Run("all flips true", func(t *testing.T) {
		random := newRandomizer(staticTime)
		f := func() bool {
			return true
		}
		random.flipper = &f
		for i := 0; i < 10000; i++ {
			t1, t2 := randTime(), randTime()
			log.Printf("now: %s, event: %s, epithet: %s",
				t1.Format(time.RFC3339), t2.Format(time.RFC3339), epithetFunction(random, t1, t2, "New York"))
		}
	})
	t.Run("all flips random", func(t *testing.T) {
		random := newRandomizer(staticTime)
		for i := 0; i < 10000; i++ {
			t1, t2 := randTime(), randTime()
			log.Printf("now: %s, event: %s, epithet: %s",
				t1.Format(time.RFC3339), t2.Format(time.RFC3339), epithetFunction(random, t1, t2, "New York"))
		}
	})
}

func randTime() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
}
