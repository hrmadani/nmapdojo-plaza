package main

///////////////TASKS//////////////////
//()Shwo all alive reports
//()Produce message to the RabbitMQ
/////////////////////////////////////
import (
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

}
