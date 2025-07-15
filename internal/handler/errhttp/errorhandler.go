package errhttp

import (
	"errors"
	"net/http"

	"github.com/Glaz97/twelvefactorapp/pkg/types"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
)

func Handler(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		var verr validation.Errors
		switch {
		case errors.Is(err, types.ErrValidation):
			abortWithStatusJSON(c, http.StatusBadRequest, err.Error())
		case errors.As(err, &validation.ErrInInvalid):
			abortWithStatusJSON(c, http.StatusBadRequest, err.Error())
		case errors.As(err, &verr):
			abortWithStatusJSON(c, http.StatusBadRequest, err.Error())
		case err.IsType(gin.ErrorTypeBind):
			abortWithStatusJSON(c, http.StatusBadRequest, err.Error())
		case errors.Is(err, types.ErrNotFound):
			abortWithStatusJSON(c, http.StatusNotFound, err.Error())
		case errors.Is(err, types.ErrConflict):
			abortWithStatusJSON(c, http.StatusConflict, err.Error())
		case errors.Is(err, types.ErrUnauthorized):
			abortWithStatusJSON(c, http.StatusUnauthorized, err.Error())
		case errors.Is(err, types.ErrForbidden):
			abortWithStatusJSON(c, http.StatusForbidden, err.Error())
		case errors.Is(err, types.ErrInternal):
			log.Error("internal error", zap.Error(err))
			abortWithStatusJSON(c, http.StatusInternalServerError, "internal error")

		default:
			log.Error("unhandled internal error", zap.Error(err))
			abortWithStatusJSON(c, http.StatusInternalServerError, "internal error")
		}
	}
}

func AbortWithError(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Abort()
}

func abortWithStatusJSON(c *gin.Context, code int, message string) {
	if c.Writer.Written() {
		return
	}
	c.AbortWithStatusJSON(
		code,
		types.ErrorJSONResponse{
			Error: types.Error{
				Code:    code,
				Message: message,
			},
		})
}
