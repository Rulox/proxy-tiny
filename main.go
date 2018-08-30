package main

import (
	"fmt"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Let's assume we have this (simplified) definition for the reverse proxy upstreams configuration
type Upstreams struct {
	RelativePath string
	Host         string
	Port         int
}

var upstreams = []Upstreams{
	{
		RelativePath: "/GET/200",
		Host:         "http_test_server", // or localhost if not running with Docker
		Port:         9980,
	},
	{
		RelativePath: "/valyala/fasthttp/issues/348",
		Host:         "github.com",
		Port:         80,
	},
	// etc
}

// HttpProxy is a struct that stores the information for each proxy
type HttpProxy struct {
	proxy       *fasthttp.HostClient
	redirectURI *fasthttp.URI
}

// Basic handler for the proxy
func (h *HttpProxy) reverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response

	// Here we do some things for Auth (check and add headers basically) [..]

	if err := h.proxy.Do(req, resp); err != nil {
		resp.SetStatusCode(fasthttp.StatusServiceUnavailable)
		fmt.Printf("error when proxying the request: %s", err)
	}
}

// InitHttpProxy initializes a basic proxy using the configured routes in `upstreams`
func InitHttpProxy() {
	router := fasthttprouter.New()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false
	router.HandleOPTIONS = false

	s := &fasthttp.Server{
		Handler:          router.Handler,
		Concurrency:      fasthttp.DefaultConcurrency,
		DisableKeepalive: false,
	}

	// Let's add all the upstreams to our fasthttp router
	for _, upstream := range upstreams {
		rPath := upstream.RelativePath

		newProxy := fasthttp.HostClient{
			Addr: upstream.Host + ":" + strconv.Itoa(upstream.Port),
		}

		proxyURI := fasthttp.URI{}
		proxyURI.SetPath(upstream.RelativePath)

		httpProxy := &HttpProxy{
			proxy:       &newProxy,
			redirectURI: &proxyURI,
		}
		fmt.Println(
			fmt.Sprintf("New Http proxy registered to %s:%d%s", upstream.Host, upstream.Port, upstream.RelativePath))
		var allowedMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "HEAD", "OPTIONS"}
		for _, method := range allowedMethods {
			router.Handle(method, rPath, httpProxy.reverseProxyHandler)
		}
	}


	fmt.Println(fmt.Sprintf("Listening for HTTP connections on port 8000"))
	s.ListenAndServe(":" + strconv.Itoa(8000))
}

func main() {
	InitHttpProxy()
}
