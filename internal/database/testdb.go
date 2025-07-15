package database

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

func NewTestDatabase(ctx context.Context, tb testing.TB, log *zap.Logger) *Database {
	tb.Helper()
	if testing.Short() {
		tb.Skip("skipping test in short mode.")
	}

	testDBName := "test_db_" + strconv.Itoa(int(rand.Uint32())) //nolint:gosec
	testCfg := config.DefaultConfig()
	testCfg.MongoDB.Database = testDBName

	uri, set := os.LookupEnv("MONGODB_URI")
	if set {
		require.NotContains(tb, uri, "mongodb.net", "Do not use production database for tests")
		testCfg.MongoDB.URI = config.SecretString(uri)
	}

	var testDB *Database
	err := fxtest.New(
		tb,
		Module,
		fx.Provide(func() *config.MongoDB { return &testCfg.MongoDB }),
		fx.Provide(func() *zap.Logger { return log }),
		fx.Invoke(func(db *Database) { testDB = db }),
	).Start(ctx)
	require.NoError(tb, err)
	require.NoError(tb, testDB.Client().Ping(ctx, nil))

	tb.Cleanup(func() {
		// To assure that parallel runs do not interfere in a cleanup
		err1 := testDB.Client().Ping(ctx, nil)
		err2 := testDB.Drop(ctx)
		err3 := testDB.Stop(ctx)
		require.NoError(tb, err1)
		require.NoError(tb, err2)
		require.NoError(tb, err3)
	})

	return testDB
}
