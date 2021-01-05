package routes

import (
	"github.com/go-chi/chi"
	"github.com/uandersonricardo/jobbing-jobs-service/controllers"
)

func star(r chi.Router) {
	r.Post("/", controllers.StarCreate)
	r.Put("/{starID}", controllers.StarUpdate)
	r.Delete("/{starID}", controllers.StarDelete)
}
