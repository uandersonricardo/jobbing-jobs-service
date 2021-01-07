package models

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Vacancies        int                  `json:"vacancies" bson:"vacancies"`
	Stars            []primitive.ObjectID `json:"stars" bson:"stars"`
	Company          primitive.ObjectID   `json:"company" bson:"company"`
	Applicants       []primitive.ObjectID `json:"applicants" bson:"applicants"`
	City             string               `json:"city" bson:"city"`
	State            string               `json:"state" bson:"state"`
	Salary           float32              `json:"salary" bson:"salary"`
	ExpiresIn        time.Time            `json:"expires_in" bson:"expires_in"`
}

func (f *Job) Creating() error {
	if f.Stars == nil {
		f.Stars = []primitive.ObjectID{}
	}

	if f.Applicants == nil {
		f.Applicants = []primitive.ObjectID{}
	}

	return nil
}

func (f *Job) Saving() error {
	if f.Stars == nil {
		f.Stars = []primitive.ObjectID{}
	}

	if f.Applicants == nil {
		f.Applicants = []primitive.ObjectID{}
	}

	return nil
}
