package service

import "net/http"

/**
 * Derived from http://thenewstack.io/make-a-restful-json-api-go/
 */

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Info",
		"GET",
		"/info",
		Info,
	},
	Route{
		"Health",
		"POST",
		"/health",
		Health,
	},
	Route{
		"Fruits",
		"GET",
		"/api/fruits",
		Fruits,
	},
	Route{
		"SwaggerJson",
		"GET",
		"/v2/api-docs",
		GetSwaggerJson,
	},
}
