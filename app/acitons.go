package app

import (
	"fmt"
	restful "github.com/emicklei/go-restful"
)

type Haproxy struct{}

func (h Haproxy) WebService() *restful.WebService {
	 ws := new(restful.WebService)

	 ws.Path("/haproxy").Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)

	 ws.Route((ws.GET("/add").To(h.add).Doc("add an item for haproxy").Param(ws.BodyParameter("name", "get name from body"))))

	 return ws
}

func (h Haproxy) add(request *restful.Request, response *restful.Response){
	fmt.Println(request)
	name := request.PathParameter("name")
	fmt.Println(name)
	response.WriteEntity("test")

}

//func (h Haproxy) list(request *restful.Request, response *restful.Response){
//	fmt.Println(request)
//	fmt.Println(response)
//}


