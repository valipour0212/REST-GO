package user

import "time"

type User struct {
	ID            string    `bson:"_id,omitempty"`
	FirstName     string    `bson:"FirstName,omitempty"`
	LastName      string    `bson:"LastName,omitempty"`
	Email         string    `bson:"Email,omitempty"`
	UserName      string    `bson:"UserName,omitempty"`
	Password      string    `bson:"Password,omitempty"`
	RegisterDate  time.Time `bson:"RegisterDate,omitempty"`
	CreatorUserID string    `bson:"CreatorUserID,omitempty"`
}
