package main

import "github.com/fedosb/golang-url-shortener/internal/controllers"

func main() {

	c := controllers.New()

	err := c.Router.Run()
	if err != nil {
		panic(err)
	}

}
