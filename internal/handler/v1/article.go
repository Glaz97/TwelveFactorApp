package v1

import (
	"net/http"

	"github.com/Glaz97/twelvefactorapp/internal/article"
	"github.com/Glaz97/twelvefactorapp/internal/database"
	"github.com/Glaz97/twelvefactorapp/internal/handler/errhttp"
	"github.com/Glaz97/twelvefactorapp/internal/marshal"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleHandler struct {
	articleService *article.Service
	log            *zap.Logger
}

func NewArticleHandler(
	db *database.Database,
	log *zap.Logger,
) *ArticleHandler {
	return &ArticleHandler{
		articleService: article.NewArticleService(db, log),
		log:            log,
	}
}

func (h *ArticleHandler) RegisterRouters(r gin.IRouter) {
	r.GET("/article/:id", h.GetArticle)
	r.POST("/article", h.CreateArticle)
}

// GetArticle godoc
//
//	@Summary	Get Article
//	@Tags		Article
//	@Param		id	path		string	true	"Article"
//	@Success	200	{object}	article.ArticleGet
//	@Router		/article/{id} [get]
func (h *ArticleHandler) GetArticle(c *gin.Context) {
	articleID, err := marshal.ObjectIDFromParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		errhttp.AbortWithError(c, err)
		return
	}

	res, err := h.articleService.GetArticle(c, articleID)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		errhttp.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateArticle godoc
//
//	@Summary	Create Article
//	@Tags		Article
//	@Param		request	body		article.ArticleCreate	true	"Article"
//	@Success	201		{object}	article.ArticleGet
//	@Router		/article [post]
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article article.ArticleCreate
	if err := marshal.BindJSON(c, &article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		errhttp.AbortWithError(c, err)
		return
	}

	res, err := h.articleService.CreateArticle(c, &article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		errhttp.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}
