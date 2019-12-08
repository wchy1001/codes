package main

import (
	restful "github.com/emicklei/go-restful"
	"github.com/wchy1001/haproxy-api/app"
	"net/http"
)


func main(){
	h := app.Haproxy{}
	ws := h.WebService()
	wsContainer := restful.NewContainer()
	wsContainer.Add(ws)
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	defer server.Close()
	server.ListenAndServe()
}