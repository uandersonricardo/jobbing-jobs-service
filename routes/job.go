package routes

import (
	"github.com/go-chi/chi"
	"github.com/uandersonricardo/jobbing-jobs-service/controllers"
)

func job(r chi.Router) {
	r.Get("/", controllers.JobList)
	r.Post("/", controllers.JobCreate)
}
