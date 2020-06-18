package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type racerFunc func(u1, u2 string) (string, error)

func TestRacer(t *testing.T) {
	t.Run("serial racer", func(t *testing.T) {
		testRacer(t, SerialRacer)
	})

	t.Run("parallel racer", func(t *testing.T) {
		testRacer(t, ParallelRacer)
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer func() { serverA.Close(); serverB.Close() }()

		_, err := ParallelRacer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := ParallelRacer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableParallelRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func testRacer(t *testing.T, racer racerFunc) {
	t.Helper()

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer func() { slowServer.Close(); fastServer.Close() }()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, _ := racer(fastURL, slowURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
