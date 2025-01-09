package dto

import "net/http"

type Task struct {
	Name    string      `json:"name"`
	Url     string      `json:"url"`
	Headers http.Header `json:"headers"`
}

type Hotel struct{}
