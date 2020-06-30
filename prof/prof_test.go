package prof

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestHoge(t *testing.T) {
	t.Skip()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	go func() {
		time.Sleep(2 * time.Second)
	}()

	time.Sleep(1 * time.Second)
}
