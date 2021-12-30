package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeRequest(method string, route string, body io.Reader) (httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, route, body)

	rr := httptest.NewRecorder()

	if err != nil {
		return *rr, err
	}
	handler := http.HandlerFunc(homeLink)
	handler.ServeHTTP(rr, req)

	return *rr, nil
}

func Test_createTodo(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTodo(tt.args.w, tt.args.r)
		})
	}
}

func Test_homeLink(t *testing.T) {
	type args struct {
		method  string
		expCode int
		expBody string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				method:  "GET",
				expCode: http.StatusOK,
				expBody: "Welcome Home!!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reponse, err := makeRequest(tt.args.method, "/", nil)

			if err != nil {
				t.Fatal(err)
			}

			if status := reponse.Code; status != tt.args.expCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.args.expCode)
			}
			if reponse.Body.String() != tt.args.expBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					reponse.Body.String(), tt.args.expBody)
			}
		})
	}
}
