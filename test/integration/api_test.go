package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Glaz97/twelvefactorapp/internal/article"
	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
	"github.com/test-go/testify/require"
)

func TestApi(t *testing.T) {
	c := resty.New()
	c.SetBaseURL("http://localhost:8056")
	c.SetHeader("Content-Type", "application/json")

	// Create Article
	newArticleTitle := "article-test-" + lo.RandomString(4, lo.AlphanumericCharset)
	resp, err := c.R().
		SetBody(fmt.Sprintf(`{"title": "%s"}`, newArticleTitle)).
		Post("/article")
	require.NoError(t, err)
	require.False(t, resp.IsError(), resp.String())
	require.Equal(t, http.StatusCreated, resp.StatusCode())
	var newArticle article.Article
	require.NoError(t, json.Unmarshal(resp.Body(), &newArticle))
	require.False(t, newArticle.ID.IsZero())
	require.Equal(t, newArticleTitle, newArticle.Title)
	require.NotEmpty(t, newArticle.CreatedAt)

	// Get Article
	resp, err = c.R().
		Get(fmt.Sprintf("/article/%s", newArticle.ID.Hex()))
	require.NoError(t, err)
	require.False(t, resp.IsError(), resp.String())
	require.Equal(t, http.StatusOK, resp.StatusCode())
	var article article.Article
	require.NoError(t, json.Unmarshal(resp.Body(), &article))
	require.Equal(t, newArticle.ID, article.ID)
	require.Equal(t, newArticle.Title, article.Title)
	require.NotEmpty(t, article.CreatedAt)
}
