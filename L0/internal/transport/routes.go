package transport

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Routes *mux.Router
}

var CacheHe Cache

func InitRouter(cache *Cache) *Router {
	CacheHe.store = cache.store
	//Create router
	router := &Router{}
	router.Routes = mux.NewRouter()

	//Create routes
	router.Routes.HandleFunc("/", Home)

	return router
}
