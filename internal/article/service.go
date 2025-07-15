package article

import (
	"context"
	"errors"
	"fmt"

	"github.com/Glaz97/twelvefactorapp/internal/database"
	"github.com/Glaz97/twelvefactorapp/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	log *zap.Logger

	articles *mongo.Collection
}

func NewArticleService(
	db *database.Database,
	log *zap.Logger,
) *Service {
	return &Service{
		articles: db.Collection(database.ArticlesCollection),
		log:      log,
	}
}

func (s *Service) Start(_ context.Context) error {
	return nil
}

func (s *Service) Stop(_ context.Context) error {
	return nil
}

func (s *Service) GetArticle(ctx context.Context, idFilter primitive.ObjectID) (*ArticleGet, error) {
	query := bson.M{}
	if !idFilter.IsZero() {
		query["_id"] = idFilter
	}

	result := s.articles.FindOne(ctx, query)
	if err := result.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, types.NewNotFoundError("article with filter:" + idFilter.Hex() + " not found")
		}
		return nil, fmt.Errorf("failed to find article: %w", err)
	}

	article := ArticleGet{}
	if err := result.Decode(&article); err != nil {
		return nil, fmt.Errorf("failed to decode article: %w", err)
	}

	return &article, nil
}

func (s *Service) CreateArticle(ctx context.Context, article *ArticleCreate) (*ArticleGet, error) {
	if article.CreatedAt == nil {
		article.CreatedAt = types.GetCurrentUTCTime()
	}

	if err := article.Validate(); err != nil {
		return nil, err
	}

	result, err := s.articles.InsertOne(ctx, article)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("article create failed, %w", database.NewDuplicateKeyError(err))
		}
		return nil, fmt.Errorf("failed to insert article: %w", err)
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, types.NewInternalError("failed to assert inserted article ID")
	}

	articleGet, err := s.GetArticle(ctx, insertedID)
	if err != nil {
		return nil, err
	}

	return articleGet, nil
}
