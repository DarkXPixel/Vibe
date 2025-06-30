package internal

import (
	"context"
	"testing"

	authpb "github.com/DarkXPixel/Vibe/proto/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeRow struct {
	pgx.Row
	scanFunc func(...interface{}) error
}

func (r *fakeRow) Scan(dest ...interface{}) error {
	return r.scanFunc(dest...)
}

func TestRegister(t *testing.T) {
	t.Setenv("JWT_SECRET", "test_secret")
	ctx := context.Background()

	mockDB := &DBPoolMock{}

	mockDB.ExecFunc = func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
		return pgconn.NewCommandTag("INSERT 0 1"), nil
	}
	mockDB.QueryRowFunc = func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
		return &fakeRow{
			scanFunc: func(dest ...interface{}) error {
				if len(dest) > 0 {
					dest[0] = "mock-user-id"
				}
				return nil
			},
		}
	}

	cfg := &Config{
		JWTSecret: "test_secret",
	}

	//mock, err := pgxmock.NewPool()
	//require.NoError(t, err)
	s := NewAuthServer(mockDB, cfg)

	req := &authpb.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := s.Register(ctx, req)

	require.NoError(t, err)
	assert.NotEmpty(t, resp.GetToken())
	assert.NotEmpty(t, resp.GetUserId())
}
