package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/swarajhoster/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main()  {

	// Router
	router := httprouter.New()

	// Controller
	uc := controllers.NewUserController(getSession())

	// Router methods
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", router)
	
}

// Connection to mongodb
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if  err != nil {
		panic(err)
	}

	return s
}