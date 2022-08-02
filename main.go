package main

import (
	v3 "DlabIDPBackend/v3"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	idp := v3.IdpWebServer{}

	// landing page functions
	idp = idp.
		StartServer().
		TestActive().
		Version()

	// macro page functions
	// fund data page functions

	idp.GinCon.Run()
}
