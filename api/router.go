package api

import "github.com/h00s/go-api-template/api/controllers"

func (api *API) setRoutes() {
	api.server.Get("/api/v1/movies", controllers.GetMoviesHandler)
	api.server.Post("/api/v1/movies", controllers.CreateMovieHandler)
}
