package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Use(MapEncoder)
	m.Get("/server/:ip", GetServerDetails)
	m.Get("/servers", GetServers)
	m.Run()
}
