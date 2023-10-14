package news

import "time"

type News struct {
	Id               string    `bson:"_id,omitempty"`
	Title            string    `bson:"Title,omitempty"`
	ShortDescription string    `bson:"ShortDescription,omitempty"`
	Description      string    `bson:"Description,omitempty"`
	ImageName        string    `bson:"ImageName,omitempty"`
	CreateDate       time.Time `bson:"CreateDate,omitempty"`
	CreatorUserId    string    `bson:"CreatorUserId,omitempty"`
}
