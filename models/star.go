package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Star struct {
	mgm.DefaultModel `bson:",inline"`
	Applicant        primitive.ObjectID `json:"applicant" bson:"applicant"`
	Job              primitive.ObjectID `json:"job" bson:"job"`
	Stars            int                `json:"stars" bson:"stars"`
}
