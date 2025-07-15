package database

import (
	"context"
	"regexp"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	"github.com/Glaz97/twelvefactorapp/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var dupKeyRegex = regexp.MustCompile(`dup key: {\s*([^:]+)`)

const (
	ArticlesCollection = "articles"
)

type Database struct {
	*mongo.Database
	log *zap.Logger
}

func NewDatabase(cfg *config.MongoDB, log *zap.Logger) (*Database, error) {
	log = log.Named("database")

	// We set default compressors so compression is enabled by default
	defaultCompressors := []string{
		"zstd",
		"zlib",
		"snappy",
	}

	opts := options.Client().SetCompressors(defaultCompressors).ApplyURI(string(cfg.URI))

	log.Info("connecting to MongoDB...", zap.String("uri", cfg.URI.Mask()))

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.Database)

	return &Database{
		Database: db,
		log:      log,
	}, nil
}

func (db *Database) Start(ctx context.Context) error {
	err := db.Client().Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Stop(ctx context.Context) error {
	return db.Client().Disconnect(ctx)
}

func NewDuplicateKeyError(err error) error {
	if matches := dupKeyRegex.FindStringSubmatch(err.Error()); len(matches) > 0 {
		return types.NewConflictError(matches[1] + " is duplicated")
	}

	return types.NewConflictError("received duplicate key error")
}
