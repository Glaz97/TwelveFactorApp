package marshal

import (
	"errors"

	"github.com/Glaz97/twelvefactorapp/pkg/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromParam(c *gin.Context, key string) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(c.Param(key))
	if err != nil {
		return primitive.NilObjectID, errors.Join(err, types.ErrValidation)
	}
	return id, nil
}

func BindJSON(c *gin.Context, v any) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return errors.Join(err, types.ErrValidation)
	}
	return nil
}
