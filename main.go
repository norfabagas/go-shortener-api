package main

import "os"

func main() {
	a := App{}

	APP_PORT := os.Getenv("PORT")
	
	a.Initialize()

	a.Run(APP_PORT)
}