package errhttp_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Glaz97/twelvefactorapp/internal/handler/errhttp"
	"github.com/Glaz97/twelvefactorapp/pkg/types"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAPIKeyValidation(t *testing.T) {
	log, err := zap.NewDevelopment()
	require.NoError(t, err)
	gin.SetMode(gin.TestMode)
	handler := errhttp.Handler(log)

	t.Run("no error", func(t *testing.T) {
		w := httptest.NewRecorder()
		t.Cleanup(func() { require.NoError(t, w.Result().Body.Close()) })
		c, _ := gin.CreateTestContext(w)

		handler(c)

		res := w.Result()
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.Nil(t, c.Errors.Last())
		require.Empty(t, c.Errors)
		require.Empty(t, w.Body.String())
		require.NoError(t, res.Body.Close())
	})

	tests := []struct {
		name            string
		err             error
		codeExpected    int
		messageExpected string
	}{
		{
			name:            "validation error",
			err:             types.ErrValidation,
			codeExpected:    http.StatusBadRequest,
			messageExpected: "validation error",
		},
		{
			name:            "ozzo validation error",
			err:             validation.Errors{"field": types.ErrValidation},
			codeExpected:    http.StatusBadRequest,
			messageExpected: "field: validation error.",
		},
		{
			name:            "not found error",
			err:             types.ErrNotFound,
			codeExpected:    http.StatusNotFound,
			messageExpected: "not found error",
		},
		{
			name:            "internal error",
			err:             types.ErrInternal,
			codeExpected:    http.StatusInternalServerError,
			messageExpected: "internal error",
		},
		{
			name:            "unknown error",
			err:             fmt.Errorf("unknown error"),
			codeExpected:    http.StatusInternalServerError,
			messageExpected: "internal error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			_ = c.Error(test.err)

			handler(c)

			res := w.Result()
			require.Error(t, c.Errors.Last())
			require.Equal(t, test.codeExpected, res.StatusCode)
			require.JSONEq(t,
				fmt.Sprintf(
					`{"error":{"code":%d,"message":"%s"}}`,
					test.codeExpected,
					test.messageExpected,
				),
				w.Body.String(),
			)
			require.NoError(t, res.Body.Close())
		})
	}
}
