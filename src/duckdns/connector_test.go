package duckdns

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetTXTRecord(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Contains(t, r.URL.RawQuery, "domains=test.domain")
		assert.Contains(t, r.URL.RawQuery, "token=test_token")
		assert.Contains(t, r.URL.RawQuery, "txt=test_txt")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	connector := NewConnector("test_token")
	connector.updateURL = ts.URL

	res, err := connector.SetTXTRecord(context.Background(), "test.domain", "test_txt")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCleanTXTRecord(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Contains(t, r.URL.RawQuery, "domains=test.domain")
		assert.Contains(t, r.URL.RawQuery, "token=test_token")
		assert.Contains(t, r.URL.RawQuery, "txt=")
		assert.Contains(t, r.URL.RawQuery, "clear=true")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	connector := NewConnector("test_token")
	connector.updateURL = ts.URL

	res, err := connector.CleanTXTRecord(context.Background(), "test.domain")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
