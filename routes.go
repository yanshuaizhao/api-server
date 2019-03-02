package main

import (
	"github.com/julienschmidt/httprouter"
	"./handles"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", handles.Index},
		Route{"DemoIndex", "GET", "/demo", handles.DemoIndex},
		Route{"DemoShow", "GET", "/demo/:id", handles.DemoShow},
		Route{"DemoCreate", "POST", "/demo", handles.DemoCreate},
	}
	return routes
}
