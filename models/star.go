package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Star struct {
	mgm.DefaultModel `bson:",inline"`
	Applicant        primitive.ObjectID `json:"applicant" bson:"applicant"`
	Stars            int                `json:"stars" bson:"stars"`
}
