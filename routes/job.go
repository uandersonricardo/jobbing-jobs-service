package routes

import (
	"github.com/go-chi/chi"
	"github.com/uandersonricardo/jobbing-jobs-service/controllers"
)

func job(r chi.Router) {
	r.Get("/", controllers.JobList)
	r.Post("/", controllers.JobCreate)
	r.Put("/{jobID}", controllers.JobUpdate)
	r.Delete("/{jobID}", controllers.JobDelete)
	r.Get("/{jobID}", controllers.JobShow)
}
