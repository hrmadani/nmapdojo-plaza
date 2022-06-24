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

var router *mux.Router

func CreateRoutes() {
	router = mux.NewRouter()
}

func InitializeRoutes() {
	plaza_routes.PlazaRoutes(router)
}

func main() {
	CreateRoutes()
	InitializeRoutes()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Plaza is up!")
}
