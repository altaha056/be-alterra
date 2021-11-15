package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.InitDB()
	// http.HandleFunc("/articles", ReadArcticles)
	// fmt.Println("server is starting")
	// http.ListenAndServe(":1234", nil)
	e:=route.New()
	e.Start(":8000")
}
