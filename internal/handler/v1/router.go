package v1

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	articleHandler *ArticleHandler
}

func NewRouter(
	articleHandler *ArticleHandler,
) *Router {
	return &Router{
		articleHandler: articleHandler,
	}
}

func (h *Router) RegisterRoutes(r gin.IRouter) {
	h.articleHandler.RegisterRouters(r)
}
