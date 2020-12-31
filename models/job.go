package models

import "github.com/kamva/mgm/v3"

type Job struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Salary           int    `json:"salary" bson:"salary"`
}
