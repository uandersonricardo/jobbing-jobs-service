package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kamva/mgm/v3"
	"github.com/uandersonricardo/jobbing-jobs-service/helpers"
	"github.com/uandersonricardo/jobbing-jobs-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

func JobList(w http.ResponseWriter, r *http.Request) {
	jobs := []models.Job{}

	err := mgm.Coll(&models.Job{}).SimpleFind(&jobs, bson.M{})

	if err != nil {
		helpers.WriteError(w, "Error decoding job")
		return
	}

	res, _ := json.Marshal(jobs)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func JobCreate(w http.ResponseWriter, r *http.Request) {
	var job models.Job

	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		helpers.WriteError(w, "Error decoding job")
		return
	}

	err = mgm.Coll(&job).Create(&job)

	if err != nil {
		helpers.WriteError(w, "Error creating job")
		return
	}

	res, _ := json.Marshal(job.ID.Hex())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func JobUpdate(w http.ResponseWriter, r *http.Request) {
	var job models.Job

	jobID := chi.URLParam(r, "jobID")
	coll := mgm.Coll(&job)

	err := coll.FindByID(jobID, &job)

	if err != nil {
		helpers.WriteError(w, "Error finding job")
		return
	}

	var updatedJob models.Job

	err = json.NewDecoder(r.Body).Decode(&updatedJob)

	if err != nil {
		helpers.WriteError(w, "Error decoding job")
		return
	}

	job.Name = updatedJob.Name
	job.Description = updatedJob.Description
	job.Vacancies = updatedJob.Vacancies
	job.City = updatedJob.City
	job.State = updatedJob.State
	job.Salary = updatedJob.Salary
	job.ExpiresIn = updatedJob.ExpiresIn

	err = mgm.Coll(&job).Update(&job)

	if err != nil {
		helpers.WriteError(w, "Error updating job")
		return
	}

	res, _ := json.Marshal(job)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func JobDelete(w http.ResponseWriter, r *http.Request) {
	var job models.Job

	jobID := chi.URLParam(r, "jobID")
	coll := mgm.Coll(&job)

	err := coll.FindByID(jobID, &job)

	if err != nil {
		helpers.WriteError(w, "Error finding job")
		return
	}

	err = mgm.Coll(&job).Delete(&job)

	if err != nil {
		helpers.WriteError(w, "Error deleting job")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(nil))
}

func JobShow(w http.ResponseWriter, r *http.Request) {
	var job models.Job

	jobID := chi.URLParam(r, "jobID")
	coll := mgm.Coll(&job)

	err := coll.FindByID(jobID, &job)

	if err != nil {
		helpers.WriteError(w, "Error finding job")
		return
	}

	res, _ := json.Marshal(job)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
