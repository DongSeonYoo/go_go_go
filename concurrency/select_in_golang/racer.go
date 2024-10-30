package selectingolang

import (
	"fmt"
	"net/http"
	"time"
)

/*
select
  - Helps you wait on multiple channels.
  - Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

httptest
  - A convenient way of creating test servers so you can have reliable and controllable tests.
  - Uses the same interfaces as the "real" net/http servers which is consistent and less for you to learn.
*/
func Racer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil

	case <-ping(b):
		return b, nil

	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
