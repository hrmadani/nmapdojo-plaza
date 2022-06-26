package main

///////////////TASKS//////////////////
//()Shwo all alive reports
//()Produce message to the RabbitMQ
/////////////////////////////////////
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	plaza_routes "github.com/hrmadani/nmapdojo-plaza/pkg/routes"
)

//Muliplexer provides
var router *mux.Router

func CreateRoutes() {
	router = mux.NewRouter()
}

//InitializeRoutes initializes the plaza's routes
func InitializeRoutes() {
	plaza_routes.PlazaRoutes(router)
}

func main() {
	CreateRoutes()
	InitializeRoutes()
	http.Handle("/", router)
	fmt.Println("Plaza is up!")
	log.Fatal(http.ListenAndServe(":8000", router))
}
