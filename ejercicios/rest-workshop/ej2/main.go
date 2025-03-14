package main

import "rest-workshop/ej2/src/router"

func main() {

	router := router.CreateRouter()

	if err := router.Run(); err != nil {
		panic(err)
	}
}
