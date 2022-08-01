package web

import (
	"net/http"
)

func RootEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Write([]byte(`{"message":"hello wordl"}`))
}
