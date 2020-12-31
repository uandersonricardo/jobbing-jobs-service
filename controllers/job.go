package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/uandersonricardo/jobbing-jobs-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

func JobList(w http.ResponseWriter, r *http.Request) {
	jobs := []models.Job{}

	err := mgm.Coll(&models.Job{}).SimpleFind(&jobs, bson.M{})

	if err != nil {
		log.Fatal("Error decoding jobs")
	}

	res, _ := json.Marshal(jobs)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func JobCreate(w http.ResponseWriter, r *http.Request) {
	var job models.Job

	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		log.Fatal("Error decoding job")
	}

	err = mgm.Coll(&job).Create(&job)

	if err != nil {
		log.Fatal("Error creating job")
	}

	res, _ := json.Marshal(job.ID.Hex())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
