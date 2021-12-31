package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type request struct {
	method string
	route  string
	body   io.Reader
}

func makeRequest(requst request, handleFunc func(http.ResponseWriter, *http.Request)) (httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(requst.method, requst.route, requst.body)

	rr := httptest.NewRecorder()

	if err != nil {
		return *rr, err
	}
	handler := http.HandlerFunc(handleFunc)
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
			reponse, err := makeRequest(request{tt.args.method, "/", nil}, homeLink)

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

func Test_getTodos(t *testing.T) {
	todo, err := json.Marshal(todos)

	if err != nil {
		t.Errorf("Something went wrong %v", err)
	}

	todoString := string(todo)

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
				expBody: todoString,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reponse, err := makeRequest(request{tt.args.method, "/todos", nil}, getTodos)

			if err != nil {
				t.Fatal(err)
			}

			// var body allTodos
			// jsonErr := json.NewDecoder(reponse.Body).Decode(&body)

			// if jsonErr != nil {
			// 	t.Fatal(jsonErr)
			// }

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
