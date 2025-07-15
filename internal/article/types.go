package article

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	CreatedAt *time.Time         `json:"createdAt" bson:"createdAt"`
}

type ArticleGet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	CreatedAt *time.Time         `json:"createdAt" bson:"createdAt"`
}

type ArticleCreate struct {
	Title     string     `json:"title" bson:"title"`
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt"`
}

func (a *ArticleCreate) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.CreatedAt, validation.Required),
	)
}
