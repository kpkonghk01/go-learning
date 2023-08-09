package racer

import "net/http"

func Racer(a, b string) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
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

// WIP: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select#write-the-minimal-amount-of-code-for-the-test-to-run-and-check-the-failing-test-output-1
