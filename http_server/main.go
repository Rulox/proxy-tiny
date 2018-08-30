package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
		"strconv"
	)

// The server will respond with a message and a status code depending on the request.
func Handler(ctx *fasthttp.RequestCtx) {
	status, _ := strconv.Atoi(ctx.UserValue("status").(string))

	ctx.Response.SetStatusCode(status)
	ctx.Response.SetBodyString(
		"Body response with method = " + string(ctx.Request.Header.Method()) + " status code = " + strconv.Itoa(status))
}


func InitHTTPServer(port int) {
	router := fasthttprouter.New()
	router.GET("/GET/:status", Handler)
	router.POST("/POST/:status", Handler)
	router.OPTIONS("/OPTIONS/:status", Handler)
	router.PUT("/PUT/:status", Handler)
	router.PATCH("/PATCH/:status", Handler)
	router.HEAD("/HEAD/:status", Handler)
	router.DELETE("/DELETE/:status", Handler)

	fmt.Println("Listening for HTTP Requests on localhost:" + strconv.Itoa(port))
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), router.Handler)
}

func main() {
	InitHTTPServer(9980)
}
