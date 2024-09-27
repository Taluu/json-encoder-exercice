package http

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Taluu/json-encoder-exercise/pkg/di"
	_ "github.com/Taluu/json-encoder-exercise/pkg/test"
)

func TestSign(t *testing.T) {
	server := di.InvokeNamed[http.Handler]("/sign")

	t.Run("failures", func(t *testing.T) {
		type test struct {
			name            string
			method          string
			body            io.Reader
			expectedCode    int
			expectedMessage string
		}

		tests := []test{
			{name: "Wrong HTTP verb", method: "GET", body: nil, expectedCode: http.StatusMethodNotAllowed, expectedMessage: "unsupported method"},
			{name: "empty body", method: "POST", body: strings.NewReader(""), expectedCode: http.StatusBadRequest, expectedMessage: "json error"},
			{name: "Not a json body", method: "POST", body: strings.NewReader("invalid json"), expectedCode: http.StatusBadRequest, expectedMessage: "json error"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				r := httptest.NewRequest(tc.method, "/sign", tc.body)
				w := httptest.NewRecorder()
				server.ServeHTTP(w, r)

				resp := w.Result()

				if resp.Header.Get("Content-type") != "application/json" {
					t.Fatalf("Expected a application/json content-type, got %s", resp.Header.Get("Content-type"))
				}

				if resp.StatusCode != tc.expectedCode {
					t.Fatalf("Did not expect HTTP %d (%s)", resp.StatusCode, resp.Status)
				}

				error := struct {
					Error string `json:"error"`
				}{}

				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&error)

				if error.Error != tc.expectedMessage {
					t.Fatalf("Did not expect message %v", error.Error)
				}
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		type test struct {
			name      string
			method    string
			data      string
			signature string
		}

		tests := []test{
			{name: "json string body", method: "GET", data: "\"a\"", signature: "baef3464e171b866453457dc8a422c534a27a1d86057dd133ceead750d9a4cc8"},
			{name: "json object body", method: "GET", data: "{\"foo\":\"bar\"}", signature: "51fb0f2895400032daf856082634c635f5fe21a2848b4b2337ebeb3fc0e9c05c"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				r := httptest.NewRequest("POST", "/sign", strings.NewReader(tc.data))
				w := httptest.NewRecorder()
				server.ServeHTTP(w, r)

				resp := w.Result()
				defer resp.Body.Close()

				if resp.Header.Get("Content-type") != "application/json" {
					t.Fatalf("Expected a application/json content-type, got %s", resp.Header.Get("Content-type"))
				}

				if resp.StatusCode != 200 {
					t.Fatalf("Did not expect HTTP %d (%s)", resp.StatusCode, resp.Status)
				}

				var gotResponse struct {
					Signature string          `json:"signature"`
					Data      json.RawMessage `json:"data"`
				}

				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&gotResponse)

				if string(gotResponse.Data) != tc.data {
					t.Fatalf("the data was not properly returned, expected %s, got %s", tc.data, gotResponse.Data)
				}

				if gotResponse.Signature != tc.signature {
					t.Fatalf("not the expected signature ; expected %s, got %s", tc.signature, gotResponse.Signature)
				}
			})
		}
	})
}
