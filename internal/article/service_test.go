package article

import (
	"context"
	"testing"

	"github.com/Glaz97/twelvefactorapp/internal/database"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestArticleService(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	log, err := zap.NewDevelopment()
	require.NoError(t, err)
	testMongoDatabase := database.NewTestDatabase(ctx, t, log)

	s := NewArticleService(
		testMongoDatabase,
		log,
	)

	newArticle1, err := s.CreateArticle(ctx, &ArticleCreate{Title: "test"})
	require.NoError(t, err)
	require.NotNil(t, newArticle1)
	require.NotNil(t, newArticle1.ID)
	require.NotNil(t, newArticle1.CreatedAt)
	require.Equal(t, newArticle1.Title, "test")

	article1, err := s.GetArticle(ctx, newArticle1.ID)
	require.NoError(t, err)
	require.Equal(t, newArticle1.Title, article1.Title)

	newArticle2, err := s.CreateArticle(ctx, &ArticleCreate{Title: "test2"})
	require.NoError(t, err)
	require.NotNil(t, newArticle2)
	require.NotNil(t, newArticle2.ID)
	require.NotNil(t, newArticle2.CreatedAt)
	require.Equal(t, newArticle2.Title, "test2")

	article2, err := s.GetArticle(ctx, newArticle2.ID)
	require.NoError(t, err)
	require.Equal(t, newArticle2.Title, article2.Title)
}
