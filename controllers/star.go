package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kamva/mgm/v3"
	"github.com/uandersonricardo/jobbing-jobs-service/helpers"
	"github.com/uandersonricardo/jobbing-jobs-service/models"
)

func StarCreate(w http.ResponseWriter, r *http.Request) {
	var star models.Star

	err := json.NewDecoder(r.Body).Decode(&star)

	if err != nil {
		helpers.WriteError(w, "Error decoding star")
		return
	}

	err = mgm.Coll(&star).Create(&star)

	if err != nil {
		helpers.WriteError(w, "Error creating star")
		return
	}

	var job models.Job

	jobID := star.Job

	err = mgm.Coll(&job).FindByID(jobID, &job)

	if err != nil {
		helpers.WriteError(w, "Error finding job")
		return
	}

	job.Stars = append(job.Stars, star.ID)

	err = mgm.Coll(&job).Update(&job)

	if err != nil {
		helpers.WriteError(w, "Error updating job")
		return
	}

	res, _ := json.Marshal(star.ID.Hex())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func StarUpdate(w http.ResponseWriter, r *http.Request) {
	var star models.Star

	starID := chi.URLParam(r, "starID")
	coll := mgm.Coll(&star)

	err := coll.FindByID(starID, &star)

	if err != nil {
		helpers.WriteError(w, "Error finding star")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&star)

	if err != nil {
		helpers.WriteError(w, "Error decoding star")
		return
	}

	err = coll.Update(&star)

	if err != nil {
		helpers.WriteError(w, "Error updating star")
		return
	}

	res, _ := json.Marshal(star)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func StarDelete(w http.ResponseWriter, r *http.Request) {
	var star models.Star

	starID := chi.URLParam(r, "starID")
	coll := mgm.Coll(&star)

	err := coll.FindByID(starID, &star)

	if err != nil {
		helpers.WriteError(w, "Error finding star")
		return
	}

	err = coll.Delete(&star)

	if err != nil {
		helpers.WriteError(w, "Error deleting star")
		return
	}

	var job models.Job

	jobID := star.Job

	err = mgm.Coll(&job).FindByID(jobID, &job)

	if err != nil {
		helpers.WriteError(w, "Error finding job")
		return
	}

	job.Stars = helpers.RemoveFromSlice(job.Stars, star.ID)

	err = mgm.Coll(&job).Update(&job)

	if err != nil {
		helpers.WriteError(w, "Error updating job")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(nil))
}

func StarShow(w http.ResponseWriter, r *http.Request) {
	var star models.Star

	starID := chi.URLParam(r, "starID")
	coll := mgm.Coll(&star)

	err := coll.FindByID(starID, &star)

	if err != nil {
		helpers.WriteError(w, "Error finding star")
		return
	}

	res, _ := json.Marshal(star)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
