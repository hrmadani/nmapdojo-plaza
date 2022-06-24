package routes

import (
	"github.com/gorilla/mux"
	"github.com/hrmadani/nmapdojo-plaza/pkg/controllers"
)

var PlazaRoutes = func(router *mux.Router) {
	router.HandleFunc("/reports", controllers.ShowAllAliveReports)
}
