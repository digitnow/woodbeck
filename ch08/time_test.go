package ch08

import (	
	"net/http"
	"time"
	"testing"
)

func TestHeadTime(t *testing.T) {
	resp, err := http.Head("https://www.time.gov")
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("Date")
	if date == "" {
		t.Fatal("ingen dato header mottatt fra time.gov")
	}

	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("time.gov: %s (avvik %s)", dt, now.Sub(dt))
}