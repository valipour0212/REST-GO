package user

import (
	"time"
)

type User struct {
	Id            string    `bson:"_id,omitempty"`
	FirstName     string    `bson:"firstName,omitempty"`
	LastName      string    `bson:"lastName,omitempty"`
	Email         string    `bson:"Email,omitempty"`
	UserName      string    `bson:"UserName,omitempty"`
	Password      string    `bson:"Password,omitempty"`
	RegisterDate  time.Time `bson:"RegisterDate,omitempty"`
	CreatorUserId string    `bson:"CreatorUserId,omitempty"`
}
