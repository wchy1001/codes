package app

import (
	//"encoding/json"
	"fmt"
	restful "github.com/emicklei/go-restful"
)

type Haproxy struct{
	Lb_name string
}

func (h Haproxy) WebService() *restful.WebService {
	 ws := new(restful.WebService)

	 ws.Path("/haproxy").Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)

	 ws.Route((ws.POST("/add").To(h.add).Doc("add an item for haproxy").Param(ws.BodyParameter("name", "get name from body"))))

	 return ws
}

func (h Haproxy) add(request *restful.Request, response *restful.Response){
	fmt.Println(request.Request.Header)
	//b, _ := json.Marshal(h)
	//fmt.Println(string(b))
	response.WriteEntity(test{18})

}

type test struct {
	Age int
}

//func (h Haproxy) list(request *restful.Request, response *restful.Response){
//	fmt.Println(request)
//	fmt.Println(response)
//}


