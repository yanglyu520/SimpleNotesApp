package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestGETPlayers_Success(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantCode int
		wantJSON map[string]string
	}{
		{
			name:     "simple success",
			path:     "/players/Yang",
			wantCode: http.StatusOK,
			wantJSON: map[string]string{
				"score": "1",
			},
		},
	}

	for _, tt := range tests {
		t.Run("returns Pepper's score", func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			response := httptest.NewRecorder()
			p.PlayerServer(response, r)

			gotBody := response.Body.String()
			gotCode := response.Code
			assert.Equal(t, tt.wantCode, gotCode)
			for k, v := range tt.wantJSON {
				assert.Equal(t, v, gjson.Get(gotBody, k).String(), "key: %s", k)
			}
		})
	}
}
