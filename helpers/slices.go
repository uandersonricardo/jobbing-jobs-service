package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IndexOfSlice(slice []primitive.ObjectID, element primitive.ObjectID) int {
	for index, value := range slice {
		if element == value {
			return index
		}
	}

	return -1
}

func RemoveFromSlice(slice []primitive.ObjectID, element primitive.ObjectID) []primitive.ObjectID {
	position := IndexOfSlice(slice, element)

	if position == -1 {
		return slice
	}

	return append(slice[:position], slice[position+1:]...)
}
