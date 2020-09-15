package main

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func SerialRacer(u1, u2 string) (string, error) {

	duration1 := measureResponseTime(u1)
	duration2 := measureResponseTime(u2)

	if duration1 > tenSecondTimeout || duration2 > tenSecondTimeout {
		return "", fmt.Errorf("timed out waiting for %s and %s", u1, u2)
	}
	if duration1 > duration2 {
		return u2, nil
	}
	return u1, nil
}

func measureResponseTime(url string) time.Duration {
	since := time.Now()
	http.Get(url)
	return time.Since(since)
}

func ParallelRacer(u1, u2 string) (string, error) {
	return ConfigurableParallelRacer(u1, u2, tenSecondTimeout)
}

func ConfigurableParallelRacer(u1, u2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", u1, u2)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
